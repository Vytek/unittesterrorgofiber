# unittesterrorgofiber
 Unit Test Error Test for issue: https://github.com/gofiber/fiber/issues/1574#issuecomment-1044061058

 develop:~/unittesterrorgofiber$ go test -v
=== RUN   TestIndexRoute
panic: interface conversion: net.Addr is fiber.testAddr, not *net.TCPAddr

goroutine 23 [running]:
github.com/gofiber/fiber/v2.(*Ctx).Port(0xc0000eab00)
        /home/gt50/go/pkg/mod/github.com/gofiber/fiber/v2@v2.27.0/ctx.go:641 +0x5c
github.com/gofiber/fiber/v2/middleware/logger.New.func2.2({0xc0000bc7f8, 0x767e6c}, {0x767e70, 0x4})
        /home/gt50/go/pkg/mod/github.com/gofiber/fiber/v2@v2.27.0/middleware/logger/logger.go:218 +0x3054
github.com/gofiber/fiber/v2/internal/fasttemplate.(*Template).ExecuteFunc(0xc0000f2a00, {0x7cba40, 0xc0000bc7f8}, 0xc0001eb770)
        /home/gt50/go/pkg/mod/github.com/gofiber/fiber/v2@v2.27.0/internal/fasttemplate/template.go:293 +0x199
github.com/gofiber/fiber/v2/middleware/logger.New.func2(0xc0000eab00)
        /home/gt50/go/pkg/mod/github.com/gofiber/fiber/v2@v2.27.0/middleware/logger/logger.go:207 +0xa6e
github.com/gofiber/fiber/v2.(*Ctx).Next(0x38)
        /home/gt50/go/pkg/mod/github.com/gofiber/fiber/v2@v2.27.0/ctx.go:789 +0x43
github.com/gofiber/fiber/v2/middleware/logger.New.func2(0xc0000eab00)
        /home/gt50/go/pkg/mod/github.com/gofiber/fiber/v2@v2.27.0/middleware/logger/logger.go:160 +0x1f7
github.com/gofiber/fiber/v2.(*App).next(0xc0001ce000, 0xc0000eab00)
        /home/gt50/go/pkg/mod/github.com/gofiber/fiber/v2@v2.27.0/router.go:132 +0x1d8
github.com/gofiber/fiber/v2.(*App).handler(0xc0001ce000, 0x47fbf7)
        /home/gt50/go/pkg/mod/github.com/gofiber/fiber/v2@v2.27.0/router.go:160 +0x45
github.com/valyala/fasthttp.(*Server).serveConn(0xc0000e6900, {0x7d4030, 0xc0000949b0})
        /home/gt50/go/pkg/mod/github.com/valyala/fasthttp@v1.33.0/server.go:2328 +0x120e
github.com/valyala/fasthttp.(*Server).ServeConn(0xc0000e6900, {0x7d4030, 0xc0000949b0})
        /home/gt50/go/pkg/mod/github.com/valyala/fasthttp@v1.33.0/server.go:1997 +0x8d
github.com/gofiber/fiber/v2.(*App).Test.func1()
        /home/gt50/go/pkg/mod/github.com/gofiber/fiber/v2@v2.27.0/app.go:908 +0x35
created by github.com/gofiber/fiber/v2.(*App).Test
        /home/gt50/go/pkg/mod/github.com/gofiber/fiber/v2@v2.27.0/app.go:907 +0x2c5
exit status 2
FAIL    github.com/unittesterrorgofiber 0.013s
T4F-Demo@algodevel:~/unittesterrorgofiber$ ./gategofiber

**Go Version used:**

go version go1.17.6 linux/amd64

