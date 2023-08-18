<?php
$host = $_ENV["MYSQL_HOST"];
$username = $_ENV["MYSQL_USER"];
$password = $_ENV["MYSQL_PASSWORD"];
$database = $_ENV["MYSQL_DATABASE"];


try {
    $pdo = new PDO("mysql:host=$host;dbname=$database", $username, $password);
    $pdo->setAttribute(PDO::ATTR_ERRMODE, PDO::ERRMODE_EXCEPTION);

    // Check if the table exists
    $tableName = 'example_table';
    $query = "SELECT 1 FROM $tableName LIMIT 1;";
    $pdo->query($query);
    
    echo "Table '$tableName' exists in the database.";
} catch (PDOException $e) {
    echo $e->getMessage();
}
