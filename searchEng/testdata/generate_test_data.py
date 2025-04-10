import pandas as pd
import numpy as np

# Create sample data
data = {
    'Message': [
        'This is a test message about technology',
        'Another message about artificial intelligence',
        'A third message discussing machine learning',
        'Message about data science and analytics',
        'Final test message about programming'
    ],
    'MessageRaw': [
        'Raw: This is a test message about technology',
        'Raw: Another message about artificial intelligence',
        'Raw: A third message discussing machine learning',
        'Raw: Message about data science and analytics',
        'Raw: Final test message about programming'
    ],
    'Tag': ['tech', 'ai', 'ml', 'ds', 'dev'],
    'Sender': ['user1', 'user2', 'user3', 'user4', 'user5'],
    'Event': ['event1', 'event2', 'event3', 'event4', 'event5'],
    'EventId': ['e1', 'e2', 'e3', 'e4', 'e5'],
    'NanoTimeStamp': [int(1e18), int(1.1e18), int(1.2e18), int(1.3e18), int(1.4e18)],
    'Namespace': ['ns1', 'ns2', 'ns3', 'ns4', 'ns5']
}

# Create DataFrame
df = pd.DataFrame(data)

# Add StructuredData column with nested data
df['StructuredData'] = [
    {'key1': 'value1', 'key2': 123},
    {'key1': 'value2', 'key2': 456},
    {'key1': 'value3', 'key2': 789},
    {'key1': 'value4', 'key2': 101},
    {'key1': 'value5', 'key2': 202}
]

# Add Groupings column with arrays
df['Groupings'] = [
    ['group1', 'group2'],
    ['group3', 'group4'],
    ['group5', 'group6'],
    ['group7', 'group8'],
    ['group9', 'group10']
]

# Save to Parquet file
df.to_parquet('test_data.parquet', index=False)
print("Test data generated successfully!") 