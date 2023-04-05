# RESTful APIs
It stands for **Representational State Transfer**.

## URIs
* **URI** stands for **Uniform Resource Identifiers** and is the method by which you will access the API.
* URL stands for **Uniform Resource Locator**. the term URL does not refer to a formal partition of URI space rather, URL is an informal concept; a URL is a type of URI that identifies a resource via its network location. 

### URI format
> URI = scheme "://" authority "/" path [ "?" query] ["#" fragment"] <br/>
> URI = http://myserver.com/mypath?query=1#document <br/>

We will use the path element in order to locate an endpoint that is running on our server.as you will use this to pass parameters such as page number or ordering to control the data that is returned.

Some general rules for URI formatting:
* A forward slash / is used to indicate a hierarchical relationship between resources
* A trailing forward slash / should not be included in URIs
* Hyphens - should be used to improve readability
* Underscores _ should **not** be used in URIs
* Lowercase letters are preferred as case sensitivity is a differentiator in the path part of a URI

#### URI path design for REST services
Paths are broken into documents, collections, stores, and controllers.

### Collections
A collection is a directory of resources typically broken by parameters to access an individual document. For example:
> GET /cats -> All cats in the <u>collection</u> <br/>
> GET /cats/1 -> Single <u>document</u> for a "cat 1"

When defining a collection, **we should always use a plural noun such as cats or people** for the collection name.

### Documents
A document is a resource pointing to a single object, similar to a row in a database. It has the ability to have child resources that may be both sub-documents or collections.
> GET /cats/1 -> Single <u>document</u> for "cat 1" <br/>
> GET /cats/1/kittens -> All kittens belonging to cat 1 <br/>
> "cat 1 = single document , kittens = collection = all kittens" <br/>
> GET /cats/1/kittens/1 -> Kitten 1 for cat 1 <br/>
> "cat 1 = single document , kitten 1 = single document"

### Controller
A controller resource is like a procedure, this is typically used when a resource cannot be mapped to standard CRUD (create,retrieve, update, and delete) functions.
* The names for controllers appear as the last segment in a URI path with no child resources.
* If the controller requires parameters,these would typically be included in the query string.
* When defining a controller name we should always use a verb. A verb is a word that indicates an action or a state of being, such as feed or send.

> POST /cats/1/feed -> Feed cat 1 <br/>
> "feed = controller" <br/>
> POST /cats/1/feed?food=fish ->Feed cat 1 a fish <br/>
> "cat 1 = document , feed = controller , food=fish = querystring as a param"

### Store
A store is a client-managed resource repository, it allows the client to add, retrieve, and delete resources. Unlike a collection, a
store will never generate a new URI it will use the one specified by the client. Take a look at the following example that would
add a new cat to our store:
> PUT /cats/2

This would add a new cat to the store with an ID of 2, if we had posted the new cat omitting the ID to a collection the response would need to include a reference to the newly defined document so we could later interact with it. Like controllers we should use a plural noun for store names.

#### HTTP verbs or CRUD function names
When designing great REST URIs we never use a CRUD function name as part of the URI, instead we use a HTTP verb. For example:
> DELETE /cats/1234

We do not include the verb in the name of the method as this is specified by the HTTP verb, the following URIs would be considered an anti-pattern:

> GET /deleteCat/1234
> DELETE /deleteCat/1234
> POST /cats/1234/delete

Each of these methods has a well-defined semantic within the context of our REST API and the correct implementation will help your user understand your intention.When we look at HTTP verbs in the next section this will make more sense.

### GET
The GET method is used to **retrieve a resource** and 
* should never be used to mutate an operation, such as updating a record.
* Typically, a body is not passed with a GET request; 
* however, it is not an invalid HTTP request to do so.

Request:
```
 GET /v1/cats HTTP/1.1
```

Response:
```
HTTP/1.1 200 OK
Content-Type: application/json
Content-Length: xxxx
{"name": "Fat Freddie's Cat", "weight": 15}
```

### POST
The POST method is used to **create a new resource** in a collection or to **execute a controller**. It is typically a non-idempotent action, in that multiple posts to create an element in a collection that will create multiple elements not updated after the first call.
The POST method is always used when calling controllers as the actions of this is considered non-idempotent.
Request:
```
POST /v1/cats HTTP/1.1
Content-Type: application/json
Content-Length: xxxx
{"name": "Felix", "weight": 5}
```
Response:
```
HTTP/1.1 201 Created
Content-Type: application/json
Content-Length: 0
Location: /v1/cats/12343
```

### PUT
The PUT method is used to **update a mutable resource** and must always **include the resource locator**. The PUT method calls are also idempotent in that multiple requests will not mutate the resource to a different state than the first call.
Request:
```
PUT /v1/cats/12343 HTTP/1.1
Content-Type: application/json
Content-Length: xxxx
{"name": "Thomas", "weight": 7 }
```
Response:
```
HTTP/1.1 201 Created/Updated
Content-Type: application/json
Content-Length: 0
```

### PATCH
The PATCH verb is used to **perform a partial update**, for example, if we only wanted to update the name of our cat we could make a PATCH request only containing the details that we would like to change.
Request:
```
PATCH /v1/cats/12343 HTTP/1.1
Content-Type: application/json
Content-Length: xxxx
{"weight": 9}
```
Response:
```
HTTP/1.1 204 No Body
Content-Type: application/json
Content-Length: 0
```
In my experience PATCH updates are rarely used, the general convention is to use a PUT and to update the whole object, this not only makes the code easier to write but also makes an API which is simpler to understand.

### DELETE
The DELETE verb is used when we want to **remove a resource**, generally we would **pass the ID** of the resource as part of the path rather than in the body of the request. This way, we have a consistent method for updating, deleting, and retrieving a document.
Request:
```
DELETE /v1/cats/12343 HTTP/1.1
Content-Type: application/json
Content-Length: 0
```
Response:
```
HTTP/1.1 204 No Body
Content-Type: application/json
Content-Length: 0
```

### HEAD
A client would use the HEAD verb when they would like to **retrieve the headers for a resource without the body**. The HEAD verb is typically used in place of a GET verb when a client only wants to check if a resource exists or to read the metadata.
Request:
```
HEAD /v1/cats/12343 HTTP/1.1
Content-Type: application/json
Content-Length: 0
```
Response:
```
HTTP/1.1 200 OK
Content-Type: application/json
Last-Modified: Wed, 25 Feb 2004 22:37:23 GMT
Content-Length: 45
```

### OPTIONS
The OPTIONS verb is used when a client would like to **retrieve the possible interactions for a resource**. Typically, the server will return an Allow header, which will include the HTTP verbs that can be used with this resource.
Request:
```
OPTIONS /v1/cats/12343 HTTP/1.1
Content-Length: 0
```
Response:
```
HTTP/1.1 200 OK
Content-Length: 0
Allow: GET, PUT, DELETE
```

#### Summary CRUD
```
GET /cats
POST /cats
GET /cats/{cID}
PUT /cats/{cID}
DELETE /cats/{cID}

HEAD /cats/{cID}
OPTIONS /cats/{cID}
```

#### URI query design
It is perfectly acceptable to use a query string as part of an API call; however, I would refrain from using this to pass data to the
service. Instead the query should be used to perform actions such as:
* Paging
* Filtering
* Sorting

If we need to make a call to a controller, we discussed earlier that we should use a POST request as this is most likely a non-idempotent request. To pass data to the service, we should include the data inside of the body. However, we could use a query

string to filter the action of the controller:
```
POST /sendStatusUpdateEmail?$group=admin
{
 "message": "Hello services lab lab lab ..."
}
```

In the preceding example, we would send a status update email with the message included in the body of the request, because
we are using the group filter passed in the query string we could restrict the action of this controller to only send to the admin
group.
If we had added the message to the query string and not passed a message body, then we would potentially be causing two
problems for ourselves. 
* The first is that the max length for a URI is 2083 characters.
* The second is that generally a POST request would always include a request body. 
Whilst this is not required by the HTTP specification, it would be expected behavior by the majority of your users.

#### Response codes
When writing a great API, we should use HTTP status codes to indicate to the client the success or failure of the request.what we will do is look at the status codes that you as a software engineer will want your microservice to return.

Bad request body:
```
POST /kittens
RESPONSE HTTP 200 OK
{
"status": 401,
"statusMessage": "Bad Request""
"errorMessage": "Name should be between 1 and 256 characters in length and only contain [A-Z] - ['-.]"'-.]"
}
```

Successful request:
```
POST /kittens
RESPONSE HTTP 201 CREATED
{
"status": 201,
"statusMessage": "Created",
"kitten": {
    "id": "1234334dffdf23",
    "name": "Freddy's Cat"
    }
}
```

### 2xx Success
2xx status codes indicate that the clients request has been successfully received and understood.

#### 200 OK
This is a generic response code indicating that the request has succeeded. The response accompanying this code is generally:
* GET: An, an entity corresponding to the requested resource
* HEAD: The, the header fields corresponding to the requested resource without the message body
* POST: An, an entity describing or containing the result of the action

#### 201 Created
The created response is sent when a request succeeds and the result is that a new entity has been created. Along with the response it is common that the API will return a Location header with the location of the newly created entity:
```
    201 Created
    Location: https://api.kittens.com/v1/kittens/123dfdf111
```
**It is optional to return an object body with this response type.**

#### 204 No Content
This status informs the client that the request has been successfully processed; however, there will be no message body with the response. For example, if the user makes a DELETE request to the collection then the response may return a 204 status.

### 3xx Redirection
The 3xx indicate class of status codes indicates that the client must take additional action to complete the request. 
* Many of these status codes are used by CDNs and other content redirection techniques, 
* however, code 304 can exceptionally useful when designing our APIs to provide semantic feedback to the client.

#### 301 Moved Permanently
This tells the client that the resource they have requested has been permanently **moved to a different location**. 

Whilst this is traditionally used to redirect a page or resource from a web server it can also be useful to us when we are building our APIs. In
the instance that **we rename a collection we could use a 301 redirect to send the client to the correct location**. This however
should be used as an exception rather than the norm. Some clients do not implicitly follow 301 redirect and implementing this
capability adds additional complexity for your consumers.301 Moved Permanently

#### 304 Not Modified
This response is generally used by a CDN or caching server and is set to indicate that the response has not been modified since

the last call to the API. This is designed to save bandwidth and the request will not return a body, but will return a Content-
Location and Expires header.

### 4xx Client Error
In the instance of an **error caused by a client, not the server**, the server will return a 4xx response and will always return an
entity that gives further details on the error.

#### 400 Bad Request
This response indicates that the request could not be understood by **the client due to a malformed request** or due to a failure of
domain validation (missing data, or an operation that would cause invalid state).

#### 401 Unauthorized
This indicates that the **request requires user authentication** and will include a WWW-Authenticate header containing a challenge
applicable to the requested resource. If the user has included the required credentials in the WWW-Authenticate header, then the
response should include an error object that may contain relevant diagnostic information.

#### 403 Forbidden
The server has understood the request, but **is refusing to fulfill it**. This could be due to **incorrect access level to a resource** not
that the user is not authenticated.
If the server does not wish to make the fact that a request is not able to access a resource due to access level public, then it is
permissible to return a 404 Not found status instead of this response.

#### 404 Not Found
This response indicates that the server has **not found anything matching the requested URI**. No indication is given of whether
the condition is temporary or permanent.
It is permissible for the client to make multiple requests to this endpoint as the state may not be permanent.

#### 405 Method Not Allowed
The method specified in the request is **not allowed for the resource indicated by the URI**. This may be when the client attempts
to mutate a collection by sending a POST, PUT, or PATCH to a collection that only serves retrieval of documents.

#### 408 Request Timeout
The client did not produce a request within the time that the server is prepared to wait. The client may repeat the request
without modification at a later time.

### 5xx Server Error
Response status codes within the 500 range indicate that something has gone "Bang", the server knows this and is sorry for the
situation.
The RFC advises that an error entity should be returned in the response explaining whether this is permanent or temporary and
containing an explanation of the error. With this in mind it is currently common that a 500 error will just return something very
generic.

#### 500 Internal Server Error
A generic error message indicating that something did not go quite as planned.

#### 503 Service Unavailable
The server is currently unavailable due to temporary overloading or maintenance. There is a rather useful pattern that you can
implement to avoid cascading failure in the instance of a malfunction in which the microservice will monitor its internal state
and in the case of failure or overloading will refuse to accept the request and immediately signal this to the client. We will look
at this pattern more in chapter xx; however, this instance is probably where you will be wanting to return a 503 status code. This
could also be used as part of your health checks.

|HTTP Verb |CRUD   | Entire Collection (e.g. /customers)  |Specific Item (e.g. /customers/{id}) |
| :------------ |:---------------| :---------------------------------------------| :-----|
| POST      | Create  | 201 (Created), 'Location' header with link to /customers/{id} containing new ID. | 404 (Not Found), 409 (Conflict) if resource already exists.. |
| GET      |Read   |  200 (OK), list of customers. Use pagination, sorting and filtering to navigate big lists. | 200 (OK), single customer. 404 (Not Found), if ID not found or invalid. |
| PUT  | Update/Replace  |   405 (Method Not Allowed), unless you want to update/replace every resource in the entire collection. | 200 (OK) or 204 (No Content). 404 (Not Found), if ID not found or invalid. |
| PATCH   | Update/Modify |   405 (Method Not Allowed), unless you want to modify the collection itself. | 200 (OK) or 204 (No Content). 404 (Not Found), if ID not found or invalid. |
| DELETE    |Delete  |  405 (Method Not Allowed), unless you want to delete the whole collection—not often desirable. | 200 (OK). 404 (Not Found), if ID not found or invalid. |


### HTTP headers
Standard request headers: <br/>
Request headers provide additional information for the request and the response of your API. Think of them like metadata for the operation.
* They can be used to augment other data for the response that does not belong in the body itself such as the content encoding.
* They can also be utilized by the client to provide information that can help the server process the response.

#### Authorization
Authorization is one of the most commonly used request headers,By requesting that the user authorizes a request, you have the capability to perform
operations such as user level logging and rate limiting.

Twitter Authorization Header : authorization type is OAuth
```
POST /1.1/statuses/update.json?include_entities=true HTTP/1.1
Accept: */*
Connection: close
User-Agent: OAuth gem v0.4.4
Content-Type: application/x-www-form-urlencoded
Authorization:
    OAuth oauth_consumer_key="xvz1evFS4wEEPTGEFPHBog",
    oauth_nonce="kYjzVBB8Y0ZFabxSWbWovY3uYSQ2pTgmZeNu2VS4cg",
    oauth_signature="tnnArxj06cWHq44gCs1OSKk%2FjLY%3D",
    oauth_signature_method="HMAC-SHA1",
    oauth_timestamp="1318622958",
    oauth_token="370773112-GmHxMAgYyLbNEtIKZeRNFsMKPR9EyMZeS9weJAEb",
    oauth_version="1.0"
Content-Length: 76
Host: api.twitter.com

status=Hello%20Ladies%20%2b%20Gentlemen%2c%20a%20signed%20OAuth%20request%21
```

#### Date
```
Date: Tue, 15 Nov 1994 08:12:31 GMT
```
The date and time that the request was sent, Timestamp of the request in RFC 3339 format.

#### Accept - content type
The requested content type for the response, such as:
* application/xml
* text/xml
* application/json
* text/javascript (for JSONP)

```
Content-Type: application/x-www-form-urlencoded
```
The content type of the body of the request (used in POST and PUT requests).

#### Content-Length
```
Content-Length: 348
```
The length of the request body in bytes

#### Accept-Encoding   (gzip, deflate)
List of acceptable encodings,The core of writing a response in a gzipped format is the compress/gzip package, which is part of the standard library. It allows
you to create a Writer interface that implements ioWriteCloser wrapping an existing io.Writer.
Look at gzip sample. [here](./gzip)

#### X-Request-ID/X-Correlation-ID
Whilst you may not directly request your clients to implement this header it may be
something that you add to requests when you call downstream services. When you are trying to debug a service that is
running in production it can be incredibly useful to be able to group all the requests by a single transaction ID. A
common practice that we will see when we look at logging and monitoring is to store all logs in a common database such
as Elastic Search. By setting the standard way of working when building many connected microservices that they pass
the correlation ID with each downstream call you will be able to query your logs in Kibana or another log query tool and
group them into a single transaction:
```
X-Request-ID: f058ebd6-02f7-4d3f-942e-904344e8cde
```