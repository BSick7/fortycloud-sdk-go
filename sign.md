# FortyCloud signatures

## Reference

http://www.jokecamp.com/blog/examples-of-creating-base64-hashes-using-hmac-sha256-in-different-languages/

## Example

Following is a java code examples that generates the Authorization data

```java
SimpleDateFormat sdf = new SimpleDateFormat(="E, dd MMM yyyy HH:mm:ss");
sdf.setTimeZone(TimeZone.getTimeZone("UTC"));
String date = sdf.format(new Date()) + " GMT";

// Set current date to Date Header
httpreq.setHeader(HTTP.DATE_HEADER, date); 

// Build string to be signed
StringBuilder stringToSign = new StringBuilder ();    

// Add HTTP  Method
stringToSign.append("POST").append("\n");       

// Add Current date (as in Date header)
stringToSign.append("application/json";).append("\n");   

// Add date (as in Date header)
stringToSign.append(date).append("\n");       

// Add request URL
stringToSign.append("/restapi/v0.4/subnets").append("\n"); 

// Add Request Body
stringToSign.append({"subnet":{"name":"TestSubnet","description":"test Subnet description","cidr":"10.10.10.0/16","disableAutoNAT":false,"actualSubnet":""}}).append("\n");  

// Calculate Signature
byte[] byteToSign = stringToSign.toString().getBytes("US-ASCII");
Mac mac = Mac.getInstance("HmacSHA256");
mac.init(new SecretKeySpec(secretKey.getBytes(), "HmacSHA256"));
byte[] rawHmac = mac.doFinal(byteToSign);
String signature = new String(Base64.encodeBase64(rawHmac));

// Set accessKey and signed request into Authoziation header
httpreq.setHeader("Authorization", "FCRestAPI AccessKey=" + accessKey + " SignatureType=HmacSHA256 Signature=" + signature);
```