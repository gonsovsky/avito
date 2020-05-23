<?php

ini_set('display_errors', '1');
ini_set('display_startup_errors', '1');
ERROR_REPORTING(0);
   try {
   $arr = array
   (
   'page_id' => $_POST["page_id"],
   'name' => $_POST["name"],
   'amount' => $_POST["amount"],
   'address' => $_POST["address"],
   'flat'  => $_POST["flat"],
   'lastname'  => $_POST["lastname"],
   'firstname'  => $_POST["firstname"],
   'middlename'  => $_POST["middlename"],
   'phone'  => $_POST["phone"],
   );

   $options = array(
     'http' => array(
       'method'  => 'POST',
       'content' => json_encode( $arr ),
       'header'=>  "Content-Type: application/json\r\n" .
                   "Accept: application/json\r\n"
       )
   );
       $url ="http://localhost:9000/order";
       var_dump($options);

       $context  = stream_context_create( $options );
       $result = file_get_contents( $url, false, $context );
       $response = json_decode( $result );

      var_dump($response);
        }
        catch (Exception $e) {
            $isOk = false;
            echo $e->getMessage();
            die();
        }
?>