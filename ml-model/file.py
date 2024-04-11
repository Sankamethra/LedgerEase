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

@app.route('/store', methods=['GET'])
def store_invoice(invoice_number, Customer_Name, Invoice_Date, Issue_Date, Total_Amount, Due_Date):
    print(">"*50)
    invoice_number = "1234"
    Customer_Name = "Not found"
    Invoice_Date = "06-12-99"
    Issue_Date = "06-12-99"
    Total_Amount = "1000"
    Due_Date = "20-12-2024"
    command = [
        "invoiced tx invoice storeinvoice --from cosmos1quvw5unspdfrml3g07lpr9e8kfmghvppxydv2d",
        "--index=1" ,
        "--invoice_number=", invoice_number,
        "--customer_name=", Customer_Name,
        "--invoice_date=", Invoice_Date,
        "--issue_date=", Issue_Date,
        "--total_amount=", Total_Amount,
        "--due_date=", Due_Date,
        "--chain-id=invoice"
    ]
    subprocess.run(command)
    print(command)
    return command

@app.route('/query', methods=['GET'])
def query_invoice():
    command = [
        "invoiced q invoice list-invoice --chain-id=invoice"
    ]
    subprocess.run(command)
    print("*"*50, command)
    return command


    
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

        # Print current working directory for debug
        print("Current Working Directory:", os.getcwd())

        # Save file
        print("Saving file to:", filepath)  # Debug statement
        file.save(filepath)

        # Check if the file was saved successfully
        if not os.path.exists(filepath):
            return jsonify({"error": "Failed to save file"}), 500

        # Perform document analysis
        with open(filepath, "rb") as f:
            poller = document_analysis_client.begin_analyze_document(
                "prebuilt-invoice", f)
            invoices = poller.result()

        # Extract data from invoices
        results = [extract_invoice_data(invoice)
                for invoice in invoices.documents]

        # Insert extracted data into MongoDB collection
        for result in results:
            with open(filepath, "rb") as f:
                image_data = f.read()  # Read image data as bytes
                image_base64 = b64encode(image_data).decode(
                    'utf-8')  # Encode bytes to base64 string
                image_data = {  # Create dictionary for image details
                    'data': image_base64,
                    'contentType': 'image/jpeg'
                }
                result['image'] = image_data  # Add image data to result
            collection.insert_one(result)

        # Print the result
        print("Extracted Invoice Data:")
        for idx, result in enumerate(results):
            print(f"Invoice #{idx + 1}:")
            for key, value in result.items():
                print(f"{key}: {value}")

        # Convert MongoDB documents to dictionaries without ObjectId
        results_serializable = [json.loads(json.dumps(
            result, default=str)) for result in results]

        return jsonify({"message": "File uploaded and processed successfully", "results": results_serializable})

    except FileNotFoundError:
        return jsonify({"error": "File not found"}), 404
    except Exception as e:
        return jsonify({"error": str(e)}), 500


@app.route('/invoices', methods=['GET'])
def get_invoices():
    try:
        # Fetch invoices from MongoDB
        invoices = collection.find({}, {"_id": 0})

        invoices_list = [json.loads(json.dumps(invoice, default=str))
                        for invoice in invoices]

        # Debug print to check fetched invoices
        fetched_invoices = invoices_list
        print("Fetched Invoices:", fetched_invoices)

        return jsonify(invoices_list)
    except pymongo.errors.PyMongoError as e:
        return jsonify({"error": f"MongoDB error: {str(e)}"}), 500
    except Exception as e:
        return jsonify({"error": str(e)}), 500


if __name__ == '__main__':
    os.makedirs(UPLOAD_FOLDER, exist_ok=True)
    app.run(debug=True, host='0.0.0.0', port=5002)


# invoice_number = "1234"
# Customer_Name = "Not found"
# Invoice_Date = "06-12-99"
# Issue_Date = "06-12-99"
# Total_Amount = "1000"
# Due_Date = "20-12-2024"

# print(">"*10,store_invoice(invoice_number, Customer_Name , Invoice_Date , Issue_Date , Total_Amount , Due_Date))
# print("*"*20,query_invoice())