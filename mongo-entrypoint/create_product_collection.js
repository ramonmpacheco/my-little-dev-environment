db = db.getSiblingDB('mylittledevenvironment');

db.createCollection('products')

db.products.insertMany([
    { name: 'IPHONE' },
    { name: 'LAPTOP I9' },
    { name: 'TV SAMSUNG' }
]);