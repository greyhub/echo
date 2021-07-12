# HTTP

1. [HTTP Headers](#httpheaders)
2. [HTTP Response Status Code](#httpstatuscode)


## HTTP Headers

**HTTP headers** là phần cốt lõi của **HTTP request** và **HTTP response**. 

HTTP headers cho phép client và server chuyển thông tin bổ sung với mỗi HTTP request và response. 

<p align="center">
  <img width="40%" height="40%" src="https://2.bp.blogspot.com/-3I7YA-5vrN0/UsLuvYJWclI/AAAAAAAABq4/L7IGV6gHH14/s320/http-header.jpg">
</p>

Theo nội dung, headers được chia thành:
+ Request headers
> chứa thêm thông tin về tài nguyên hay về client.
+ Response headers
> giữ thông tin thêm về phản hồi (như vị trí của phản hồi, hay về server).
+ Representation headers
> chứa thông tin về nội dung tài nguyên (như kiểu MIME, hoặc kiểu nén/mã hóa).
+ Payload headers
> bao gồm độ dài nội dung và mã hóa được dùng để vận chuyển (transport).

Theo cách xử lý, headers được chia thành:
+ Connection
+ Keep-Alive
+ Proxy-Authenticate
+ Proxy-Authorization
+ TE
+ Trailer
+ Transfer-Encoding
+ Upgrade 

End-to-end headers
> Là những header phải được truyền đến người nhận cuối (server đối với request, client đối với response). Bộ nhớ đệm sẽ lưu trữ chúng và các proxy trung gian phải
> truyền các header chưa được sửa đổi này.

Hop-by-hop headers
> Chỉ có ý nghĩa đối với một kết nối ở mức vận chuyển duy nhất, và không được truyền lại bởi proxy, hay lưu vào bộ nhớ đệm. Hop-by-hop chỉ có thể được thiết lập
> bằng cách sử dụng Connection header.


## HTTP Status Code

Trong mỗi HTTP response đều chứa một status code.
Các status code này cho biết kết quả server trả về.

Các status code được chia thành 5 nhóm:
+ 1__: Informational responses (Thông tin)
> Request đã được server tiếp nhận và quá trình xử lý request đang được tiếp tục.
+ 2__: Successful responses (Thành công)
> Request đã được server tiếp nhận, hiểu và xử lý thành công.
+ 3__: Redirects (Chuyển hướng)
> Client cần có thêm action để hoàn thành request.
+ 4__: Client Errors (Lỗi Client)
> Request chứa cú pháp không chính xác hoặc không được thực hiện.
+ 5__: Server Errors (Lỗi Server)
> Server thất bại với việc thực hiện một request nhìn như có vẻ khả thi.

Một số status code phổ biến:

100 Continue
> Mọi thứ vẫn OK, client nên tiếp tục yêu cầu hoặc bỏ qua.

200 OK
> Request đã được tiếp nhận và xử lý thành công. 

204 No Content
> Server đã xử lý thành công request nhưng không trả về bất cứ content nào.

301 Moved Permanently
> URL được yêu cầu đã được chuyển đổi vĩnh viễn sang URL mới.

400 Bad Request
> Cú pháp của request không hợp lệ.

401 Unauthorized
> Lỗi xác thưc, ví dụ như khi nhập sai mật khẩu.

403 Forbidden
> Truy cập không được phép.

404 Not Found
> Truy vấn không được tìm thấy.

405 Method Not Allowed
> Phương thức không được hỗ trợ.

408 Request Timeout
> Quá thời gian chờ được chỉ định của server.

500 Internal Server Error
> Trường hợp máy chủ không biết cách xử lý.

503 Service Unavailable
> Ứng dụng trên server không phản hồi lại truy vấn của client, hoặc có thể ứng dụng này không hoạt động nữa.


## Ref
[Mozilla-Header](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers)
[Mozilla-Status](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status)
[Passionery](https://passionery.blogspot.com/2014/01/http-header-la-gi.html)
[Wiki](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes)
