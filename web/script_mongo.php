<?php

ini_set('display_errors', '1');
ini_set('display_startup_errors', '1');
ERROR_REPORTING(E_ALL);
require 'vendor/autoload.php';

   try {


    $id="5ec24750cbb6ebde40c59195";

        if (in_array("REQUEST_URI", $_SERVER)) {
            $id = $_SERVER['REQUEST_URI'];
            $id = substr($id, 1);
        }

    $mongo = new MongoDB\Client("mongodb://mongo-root:passw0rd@127.0.0.1:27017");
    $collection = $mongo->db->avito;
    $bid = new MongoDB\BSON\ObjectID($id);
    $result = $collection->findOne(array('_id' => $bid));
    $view = (object) array();

    $view->Title = $result["title"];
    $view->Price = "0";
        }
        catch (Exception $e) {
            echo $e->getMessage();
            die();
        }
?>