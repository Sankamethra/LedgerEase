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

def store_invoice(invoice_number, customer_name, invoice_date, total_amount, due_date):
    print(">" * 50)
    command = [
        "invoiced",
        "tx",
        "invoice",
        "storeinvoice",
        "--from",
        "cosmos1lpcgfh9gjs7r87sqrzjes3asmamd8k009j45k6",
        "--chain-id=invoice",
        "--index=1",
        f"--invoice-number={invoice_number}",
        f"--customer-name={customer_name}",
        f"--invoice-date={invoice_date}",
        f"--total-amount={total_amount}",
        f"--due-date={due_date}"
    ]
    print("Running command:", " ".join(command))  
    subprocess.run(command)
    return command

def query_invoice():
    command = [
        "invoiced",
        "q",
        "invoice",
        "list-invoice",
        "--chain-id=invoice"
    ]
    subprocess.run(command)
    print("*"*50, command)
    return command 

invoice_number = "INV-100"
Customer_Name = "MICROSOFT CORPORATION"
Invoice_Date = "2019-11-15"
Total_Amount = "110.0 USD"
Due_Date = "2019-12-15"

print(">"*10,store_invoice(invoice_number, Customer_Name , Invoice_Date , Total_Amount , Due_Date))
print("*"*20,query_invoice())