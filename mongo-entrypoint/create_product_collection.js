const db = db.getSiblingDB('mylittledevenvironment');

db.products.insertMany([
  { name: 'IPHONE' },
  { name: 'LAPTOP I9' },
  { name: 'TV SAMSUNG' }
]);
