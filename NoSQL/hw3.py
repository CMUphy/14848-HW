import boto3
import csv

s3 = boto3.resource('s3',
 aws_access_key_id='AKIAW3YYNH7JMCFHZPSV',
 aws_secret_access_key='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'
)
try:
 s3.create_bucket(Bucket='848-hw3', CreateBucketConfiguration={
 'LocationConstraint': 'us-west-2'})
except Exception as e:
 print (e)

bucket = s3.Bucket("848-hw3")
bucket.Acl().put(ACL='public-read')

# body = open('exp1.csv', 'rb')
# o = s3.Object('848-hw3', 'test').put(Body=body )
# s3.Object('848-hw3', 'test').Acl().put(ACL='public-read')


dyndb = boto3.resource('dynamodb',
 region_name='us-west-2',
 aws_access_key_id='AKIAW3YYNH7JMCFHZPSV',
 aws_secret_access_key='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'
)

try:
    table = dyndb.create_table(
        TableName='DataTable',
        KeySchema=[
            {
             'AttributeName': 'PartitionKey',
             'KeyType': 'HASH'
             },
             {
             'AttributeName': 'RowKey',
             'KeyType': 'RANGE'
             }
        ],
         AttributeDefinitions=[
         {
                 'AttributeName': 'PartitionKey',
                 'AttributeType': 'S'
         },
         {
                 'AttributeName': 'RowKey',
                 'AttributeType': 'S'
         },

     ],
     ProvisionedThroughput={
         'ReadCapacityUnits': 5,
         'WriteCapacityUnits': 5
     }
 )
except Exception as e:
     print ("The table may already exist")
     #if there is an exception, the table may already exist. if so...
     table = dyndb.Table("DataTable")

table.meta.client.get_waiter('table_exists').wait(TableName='DataTable')

with open('experiments.csv', 'r') as csvfile:
    csvf = csv.reader(csvfile, delimiter=',', quotechar='|')
    header = next(csvf)
    for item in csvf:
            print (item)
            body = open(item[4], 'rb')
            s3.Object('848-hw3', item[4]).put(Body=body)
            md = s3.Object('848-hw3', item[4]).Acl().put(ACL='public-read')

            url = " https://s3-us-west-2.amazonaws.com/848-hw3/" + item[4]
            metadata_item = {'PartitionKey': item[4], 'RowKey': item[0],
                             'Temp': item[1], 'Conductivity': item[2],
                             'Concentration':item[3],
                             'url': url}
            try:
                table.put_item(Item=metadata_item)
            except:
                print ("item may already be there or another failure")

key1={
    'PartitionKey': 'exp1.csv',
    'RowKey': '1'
}
key2={
    'PartitionKey': 'exp2.csv',
    'RowKey': '2'
}
key3={
    'PartitionKey': 'exp3.csv',
    'RowKey': '3'
}
response1 = table.get_item(Key = key1)
response2 = table.get_item(Key = key2)
response3 = table.get_item(Key = key3)
item1 = response1['Item']
item2 = response2['Item']
item3 = response3['Item']

print("Query 1 is:" + repr(key1))
print("Query 2 is:" + repr(key2))
print("Query 3 is:" + repr(key3))

print("Query 1 result:" + repr(item1))
print("Query 2 result:" + repr(item2))
print("Query 3 result:" + repr(item3))

print("Response 1"+ repr(response1))