import sys, os, pyodbc
from bs4 import BeautifulSoup
import pandas as pd
import requests

url = 'https://nafaflyball.com/database/active_dogs.php'
response = requests.get(url)
soup = BeautifulSoup(response.text, 'html.parser')
text = soup.get_text()

s = text.splitlines()

for i in s:
    j = i.split(",")

# conn_str = (
#     r'DRIVER={Microsoft Access Driver (*.mdb, *.accdb)};'
#     r'DBQ=C:\Users\admin\Documents\Megan.accdb;'
#     )
# connection = pyodbc.connect(conn_str)


# cursor = connection.cursor()

# for i in s:
#     sql="Insert into Dogs (CRN,Dog_Name,Cert_Name,Breed,Club_No,Owner_Name,Points) values ('000001','Tazzy','','Mix','266','Cathy & Forest Beaton',22597)"
#     cursor.execute(sql)

# connection.commit() 