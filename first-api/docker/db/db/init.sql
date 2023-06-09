CREATE TABLE mydb (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255),
  email VARCHAR(255),
  phone VARCHAR(255),
  address VARCHAR(255)
);
CREATE TABLE mobiles (
  id INT AUTO_INCREMENT PRIMARY KEY,
  brand VARCHAR(255),
  model VARCHAR(255),
  price INT,
  memory INT,
  camera INT
);


USE mobiles;
