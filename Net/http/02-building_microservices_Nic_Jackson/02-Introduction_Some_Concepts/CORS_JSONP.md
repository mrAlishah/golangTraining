# Accessing APIs from JavaScript
Web browsers implement a sandbox mechanism that restricts resources in one domain from accessing resources in another. For
example, you may have an API that allows the modification and retrieval of user data and a website that provides an interface
for this API. If the browser did not implement the "same-origin policy" and assuming the user did not log out of their session
then it would be possible for a malicious page to send a request to the API and modify it without you knowing.
To get around this, there are two methods that can be implemented by your microservice to allow this access, **JSONP** which
stands for (**JSON with Padding**) and **CORS** (**Cross-Origin Resource Sharing**).

## JSONP
JSONP is pretty much a hack, and it is implemented by most browsers that do not implement the later CORS standard. It is
restricted to GET requests only and works by getting round the issue that while XMLHTTPRequest is blocked from making
requests to third-party servers, there are no restrictions on HTML script elements.
A JSONP request inserts a <script src="..."> element into the browsers DOM with the API's URI as the src target. This
component returns a function call with the JSON data as a parameter, and when this loads, the function executes passing the
data to the callback.
JavaScript callback is defined in the code:
```javascript
function success(data) {
    alert(data.message);
}
```

This is the response from the API call:
```
    success({"({"message":"":"Hello World"})"})
```
To denote a request for data to be returned as JSONP, generally the `callback=functionName` parameter is added to the URI, in
our example this would be `/helloworld?callback=success`. Implementing this is particularly straightforward let's take a look
at our simple Go helloworld example and see how we can modify this to implement JSONP.
One thing to note is the Content-Type header that we are returning. We are no longer returning application/json as we are
not returning JSON we are actually returning JavaScript so we must set the Content-Type header accordingly:
```
    Content-Type: application/javascript
```
Request:
```
    GET /helloworld?callback=hello
```
Response:
```
    hello({"message":"Hello World"})
```