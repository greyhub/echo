# Cookie

[What's cookie ?](#what) \
[Why should use cookie ?](#why) \
[When's cookie created ?](#when) \
[Where's cookie stored ?](#where) \
[How ?](#how)

## What

An HTTP cookie (web cookie, browser cookie) is a small piece of data that a server sends to the user's web browser.

## Why

**Session management**
Logins, shopping carts, game scores, or anything else the server should remember.

**Personalization**
User preferences, themes, and other settings.

**Tracking**
Recording and analyzing user behavior.

## When

<p align="center">
    <img width="100%" height="100%" src="https://images.viblo.asia/b84f08a3-4f79-40b1-96a6-7a8354d6d156.jpg">
    <i>Source: viblo</i>
</p>

Cookies are created by a web server while a user is browsing a website.

## Where

Cookies are placed on the device used to access a website, and more than one cookie may be placed on a user’s device during a session.

## How

Cookies are set using the Set-Cookie header field, sent in an HTTP response from the web server. 

As an example, the browser sends its first HTTP request for the homepage of the www.example.org website:

```
GET /index.html HTTP/1.1
Host: www.example.org
...
```

The server responds with two Set-Cookie header fields:

```
HTTP/1.0 200 OK
Content-type: text/html
Set-Cookie: theme=light
Set-Cookie: sessionToken=abc123; Expires=Wed, 09 Jun 2021 10:18:14 GMT
...
```

