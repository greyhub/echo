# Gzip

## What
> *Gzip là gì ?*

Gzip Compression là phương pháp nén làm giảm dung lượng dữ liệu ở server khi gửi đến client giúp tiết kiệm băng thông, tăng tốc độ tải của website.

## Why
> *Tại sao cần sử dụng Gzip ?*

+ Tiết kiệm băng thông
+ Tăng tốc độ tải

## When
> *Nén được thực hiện khi nào ?*

<p align="center">
  <img src="https://betterexplained.com/wp-content/webp-express/webp-images/uploads/compression/HTTP_request_compressed.png.webp">
  <br>
  <i>Source: Better Explained</i>
</p>

Sau khi tìm được file trong hệ thống, file sẽ được nén để chuẩn bị gửi tới browser.

## Where
> *Nén ở đâu ?*

Nén Gzip được thực hiện từ phía server.

## How
> *Gzip được áp dụng như thế nào ?*

[Sử dụng middleware Gzip của echo framework](https://echo.labstack.com/middleware/gzip/)

Một số lưu ý khi áp dụng Gzip:
+ Không tương thích với trình duyệt quá cũ 
+ Không nên lạm dụng việc nén, chỉ nên áp dụng đối với bộ 3 html+css+js
+ Chiếm CPU của server

## Ref

[Better Explained](https://betterexplained.com/articles/how-to-optimize-your-site-with-gzip-compression/)

[Mắt bão](https://wiki.matbao.net/gzip-la-gi-cach-them-gzip-tang-toc-do-tai-website-hieu-qua/#co-che-hoat-dong-cua-gzip-la-gi)
