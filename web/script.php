<?php

ini_set('display_errors', '1');
ini_set('display_startup_errors', '1');
ERROR_REPORTING(E_ALL);
$isOk=true;

   try {
            $id = $_SERVER['REQUEST_URI'];
            $strArray = explode('/',$id);
            $id = end($strArray);

       //var_dump($id);
       $url="http://localhost:9000/page/".$id;
       $json = file_get_contents($url);
       $view = json_decode($json);
       $view->PriceDeliver = intval($view->PriceInt) + 290;
       if (property_exists($view, "Title") ==false)
       $isOk = false;
       //var_dump($view);
        }
        catch (Exception $e) {
            $isOk = false;
            echo $e->getMessage();
            die();
        }
?>