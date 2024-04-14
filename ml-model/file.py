from flask import Flask, request, jsonify, make_response
from flask_cors import CORS
from azure.core.credentials import AzureKeyCredential
from azure.ai.formrecognizer import DocumentAnalysisClient
import os
import uuid
from datetime import datetime
from pymongo import MongoClient
from bson import Binary, ObjectId  # Import ObjectId
import json  # Import json module
from base64 import b64encode
import pymongo
import subprocess
from subprocess import check_output, STDOUT

import requests

app = Flask(__name__)
CORS(app)

# Azure Form Recognizer setup
endpoint = "https://team-5-invoice.cognitiveservices.azure.com/"
key = "aaba86ed6f474bcd829d286684a31743"
document_analysis_client = DocumentAnalysisClient(
    endpoint=endpoint, credential=AzureKeyCredential(key))

# MongoDB setup
mongo_client = MongoClient('mongodb://localhost:27017/')
db = mongo_client['ledger-ease']
collection = db['invoice']

# Store the file temporarily
UPLOAD_FOLDER = 'images'
app.config['UPLOAD_FOLDER'] = UPLOAD_FOLDER


def generate_unique_filename(filename):
    # Generate a unique filename using timestamp and uuid
    timestamp = datetime.now().strftime('%Y%m%d%H%M%S')
    unique_id = str(uuid.uuid4().hex)
    _, extension = os.path.splitext(filename)
    return f"{timestamp}_{unique_id}{extension}"


def extract_invoice_data(invoice):
    # Extract specific fields from the analyzed invoice
    return {
        "Invoice Number": invoice.fields.get("InvoiceId").value if invoice.fields.get("InvoiceId") else "Not found",
        "Customer Name": invoice.fields.get("CustomerName").value if invoice.fields.get("CustomerName") else "Not found",
        "Invoice Date": str(invoice.fields.get("InvoiceDate").value) if invoice.fields.get("InvoiceDate") else "Not found",
        "Issue Date": str(invoice.fields.get("InvoiceDate").value) if invoice.fields.get("InvoiceDate") else "Not found",
        "Total Amount": f"{invoice.fields.get('InvoiceTotal').value.amount} {invoice.fields.get('InvoiceTotal').value.code}" if invoice.fields.get("InvoiceTotal") else "Not found",
        "Due Date": str(invoice.fields.get("DueDate").value) if invoice.fields.get("DueDate") else "Not found",
    }

@app.route('/store', methods=['POST'])
def store_invoice():
    try:
        # Fetch the latest index from the chain
        index_command = ["invoiced", "q", "invoice", "list-invoice", "--chain-id=invoice"]
        index_output = subprocess.run(index_command, capture_output=True, text=True)
        if index_output.returncode != 0:
            return jsonify({"error": "Failed to fetch the latest index"}), 500

        # Parse the output to extract the latest index
        latest_index = None
        try:
            output_lines = index_output.stdout.strip().splitlines()
            index_values = [int(line.split(":")[-1].strip().strip('"')) for line in output_lines if "index:" in line]
            print("Index values", index_values)
            if index_values:
                latest_index = max(index_values) + 1
                print("Latest Index (Before Increment):", latest_index)  # Debug statement
            else:
                return jsonify({"error": "No index values found in the output"}), 500
        except Exception as e:
            return jsonify({"error": f"Failed to parse index output: {str(e)}"}), 500

        data = request.get_json()  # Get JSON data from the request
        if not data:
            return make_response(jsonify({"error": "No JSON data received"}), 400)

        # Extract data from JSON
        invoice_number = data.get("Invoice Number", "")
        customer_name = data.get("Customer Name", "")
        invoice_date = data.get("Invoice Date", "")
        total_amount = data.get("Total Amount", "")
        due_date = data.get("Due Date", "")

        # Run the store_invoice command with the incremented index
        store_command = [
            "invoiced",
            "tx",
            "invoice",
            "storeinvoice",
            "--from",
            "cosmos1quvw5unspdfrml3g07lpr9e8kfmghvppxydv2d",
            "--chain-id=invoice",
            f"--index={latest_index}",
            f"--invoice-number={invoice_number}",
            f"--customer-name={customer_name}",
            f"--invoice-date={invoice_date}",
            f"--total-amount={total_amount}",
            f"--due-date={due_date}"
        ]
        subprocess.run(store_command)
        print("Running Query command:", " ".join(store_command))

        return jsonify({"message": "Invoice stored successfully", "index": latest_index})

    except Exception as e:
        return jsonify({"error": str(e)}), 500

@app.route('/query', methods=['GET'])
def query_invoice():
    try:
        # Fetch invoices from blockchain using invoiced command
        command_output = check_output(["invoiced", "q", "invoice", "list-invoice", "--chain-id=invoice"], stderr=STDOUT, text=True)
        invoices_list = parse_command_output(command_output)
        
        # Ensure invoices_list is a JSON serializable object (list of dictionaries)
        if invoices_list is None:
            return jsonify({"error": "Failed to parse invoice list"}), 500
                
        return jsonify({"message": "Invoice queried successfully", "invoice": invoices_list})
    except subprocess.CalledProcessError as e:
        # Handle subprocess errors
        return jsonify({"error": f"Command execution error: {str(e)}"}), 500
    except Exception as e:
        return jsonify({"error": str(e)}), 500
    
def parse_command_output(output):
    try:
        invoices = []
        lines = output.strip().split('\n')
        current_invoice = {}  # Initialize an empty dictionary for the current invoice
        for line in lines:
            if line.startswith('- ') and current_invoice:  # Check if a new invoice entry is starting
                if "voice" not in current_invoice:
                    invoices.append(current_invoice)  # Append the current invoice to the list of invoices if it's not a voice entry
                current_invoice = {}  # Reset the current invoice for the new entry
            key_value = line[2:].split(': ', 1)  # Remove the leading "- " and split each line into key and value
            key = key_value[0].strip()
            value = key_value[1].strip() if len(key_value) > 1 else ""  # Handle cases where value is missing
            current_invoice[key] = value
        if current_invoice and "voice" not in current_invoice:  # Append the last invoice if it exists and is not a voice entry
            invoices.append(current_invoice)
        return invoices
    except Exception as e:
        return None



def process_invoice_entry(entry):
    try:
        entry_parts = entry.strip().split('\n')  # Split entry into lines and remove leading/trailing spaces
        invoice = {}
        skip_entry = False  # Flag to determine if entry should be skipped
        for part in entry_parts:
            key_value = part.split(': ', 1)  # Split each line into key and value (max split 1 to avoid issues with colons in values)
            key = key_value[0].strip()
            value = key_value[1].strip() if len(key_value) > 1 else ""  # Handle cases where value is missing
            
            # Check if the entry should be skipped
            if "{voice:: ''}" in value or key == "creator":
                skip_entry = True
                break  # Skip this entry
            
            invoice[key] = value
        
        if not skip_entry:
            return invoice
        else:
            return None  # Return None for skipped entries
    except Exception as e:
        return None


    
@app.route('/upload', methods=['POST'])
def upload_file():
    try:
        if 'file' not in request.files:
            return make_response(jsonify({"error": "No file part"}), 400)

        file = request.files['file']
        if file.filename == '':
            return make_response(jsonify({"error": "No selected file"}), 400)

        # Generate a unique filename
        unique_filename = generate_unique_filename(file.filename)
        filepath = os.path.join(app.config['UPLOAD_FOLDER'], unique_filename)

        # Save file
        file.save(filepath)

        # Check if the file was saved successfully
        if not os.path.exists(filepath):
            return jsonify({"error": "Failed to save file"}), 500

        # Store the image file in MongoDB
        with open(filepath, "rb") as f:
            image_data = f.read()  # Read image data as bytes
            image_base64 = b64encode(image_data).decode('utf-8')  # Encode bytes to base64 string
            image_doc = {
                'filename': unique_filename,
                'data': image_base64,
                'contentType': 'image/jpeg'
            }
            collection.insert_one(image_doc)

        # Perform document analysis
        with open(filepath, "rb") as f:
            poller = document_analysis_client.begin_analyze_document("prebuilt-invoice", f)
            invoices = poller.result()

        # Extract data from invoices
        results = [extract_invoice_data(invoice) for invoice in invoices.documents]

        # Call the store_invoice function with extracted data
        for result in results:
            response = requests.post('http://localhost:5002/store', json=result)
            if response.status_code != 200:
                print("Failed to store invoice:", response.json())

        # Return the image filename
        return jsonify({"message": "File uploaded and processed successfully", "image_filename": unique_filename})

    except FileNotFoundError:
        return jsonify({"error": "File not found"}), 404
    except Exception as e:
        return jsonify({"error": str(e)}), 500

def generate_unique_filename(filename):
    # Generate a unique filename using timestamp and uuid
    timestamp = datetime.now().strftime('%Y%m%d%H%M%S')
    unique_id = str(uuid.uuid4().hex)
    _, extension = os.path.splitext(filename)
    return f"{timestamp}_{unique_id}{extension}"

@app.route('/invoices', methods=['GET'])
def get_invoices():
    try:
        # Fetch invoices from MongoDB
        invoices = collection.find({}, {"_id": 0})

        invoices_list = [json.loads(json.dumps(invoice, default=str))
                        for invoice in invoices]

        # Debug print to check fetched invoices
        fetched_invoices = invoices_list
        # print("Fetched Invoices:", fetched_invoices)

        return jsonify(invoices_list)
    except pymongo.errors.PyMongoError as e:
        return jsonify({"error": f"MongoDB error: {str(e)}"}), 500
    except Exception as e:
        return jsonify({"error": str(e)}), 500

@app.route('/invoices/<invoice_id>/image', methods=['GET'])
def get_invoice_image(invoice_id):
    try:
        # Retrieve the invoice document from MongoDB based on the invoice ID
        invoice_document = collection.find_one({"_id": ObjectId(invoice_id)})
        if invoice_document:
            image_base64 = invoice_document.get('data', '')  # Assuming 'data' contains the base64 encoded image
            return jsonify({"data": image_base64})
        else:
            return jsonify({"error": "Invoice not found"}), 404
    except pymongo.errors.PyMongoError as e:
        return jsonify({"error": f"MongoDB error: {str(e)}"}), 500
    except Exception as e:
        return jsonify({"error": str(e)}), 500


if __name__ == '__main__':
    os.makedirs(UPLOAD_FOLDER, exist_ok=True)
    app.run(debug=True, host='0.0.0.0', port=5002)

