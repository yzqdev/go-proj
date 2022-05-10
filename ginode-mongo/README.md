# ginode-mongo
一个和nodejs express很像的框架

## 本地开发
> go run main.go

beego框架自带了本地调试工具，在修改代码后可以自动重启，幸运的是，`bee`工具同样可以在`gin`项目中使用
> bee run

但是`bee`不能使用在非`GOPATH`目录下

另外一种使用了`dogo`，本项目包含了dogo.json配置，可根据自己项目路径修改
> go get github.com/liudng/dogo

> dogo

## 目录结构
本项目结构组织使用的是按职责划分，采用这种结构的原因是，目前流行的开发模式是微服务架构，但是一般项目都是由小到大再到拆分的过程，如果项目初始就使用微服务的架构开发的话估计还没等项目开发完公司就完蛋了。
所以最开始还是单体架构才是正确的方式，不过为了以后方便拆分，可以对项目目录进行合理的划分。
从路由入口看：
```golang
func RouterRegister(router *gin.Engine) {
	//user路由注册
    	userRouter := router.Group("/user")
    	{
    		userRouter.GET("/index", user.Index)
    		userRouter.GET("/users/:id", middleware.Auth(), user.GetUserInfo)
    		userRouter.GET("/users-by-sql/:id", user.GetUserInfoBySql)
    		userRouter.POST("/", user.CreateUser)
    		userRouter.PUT("/update/:id", user.UpdateUser)
    		userRouter.DELETE("/del/:id", user.DeleteUser)
    	}
    	//user路由注册
    	postRouter := router.Group("/post")
    	{
    		postRouter.GET("/", post.Index)
    		postRouter.GET("/posts/:id", middleware.Auth(), post.GetPostInfo)
    		postRouter.GET("/posts-by-sql/:id", post.GetPostInfoBySql)
    		postRouter.POST("/", post.CreatePost)
    		postRouter.PUT("/:id", post.UpdatePost)
    		postRouter.DELETE("/:id", post.DeletePost)
    	}
}
```
本项目只是一个示例项目，模块较少，只开设了两个模块，`user`和`post`，注册路由的时候就可以按照不同的职责来分组注册，`pkg`目录下放置的就是不同职责的模块。
将来需要对服务进行拆分的时候只需要将`pkg`目录下的模块分离出去即可形成独立的服务，对依赖的的修改也较少。如果使用的MVC架构的话，要对不同职责的服务进行拆分则很困难，需要对每一层中对应的文件都拿出来重新组织。

### 请求参数绑定
推荐使用**ShouldBind***，因为**Bind***是必须绑定，所有参数需要一致。
```shell script
golang.org/x/text => github.com/golang/text latest

 golang.org/x/net => github.com/golang/net latest

 golang.org/x/crypto => github.com/golang/crypto latest

 golang.org/x/tools => github.com/golang/tools latest

 golang.org/x/sync => github.com/golang/sync latest

 golang.org/x/sys => github.com/golang/sys latest

 cloud.google.com/go => github.com/googleapis/google-cloud-go latest

 google.golang.org/genproto => github.com/google/go-genproto latest

 golang.org/x/exp => github.com/golang/exp latest

 golang.org/x/time => github.com/golang/time latest

 golang.org/x/oauth2 => github.com/golang/oauth2 latest

 golang.org/x/lint => github.com/golang/lint latest

 google.golang.org/grpc => github.com/grpc/grpc-go latest

 google.golang.org/api => github.com/googleapis/google-api-go-client latest

 google.golang.org/appengine => github.com/golang/appengine latest

 golang.org/x/mobile => github.com/golang/mobile latest

 golang.org/x/image => github.com/golang/image latest
 
 cloud.google.com/go => github.com/googleapis/google-cloud-go v0.34.0

 github.com/go-tomb/tomb => gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7

 go.opencensus.io => github.com/census-instrumentation/opencensus-go v0.19.0

 go.uber.org/atomic => github.com/uber-go/atomic v1.3.2

 go.uber.org/multierr => github.com/uber-go/multierr v1.1.0

 go.uber.org/zap => github.com/uber-go/zap v1.9.1
 
 google.golang.org/api => github.com/googleapis/google-api-go-client v0.0.0-20181220000619-583d854617af

 google.golang.org/appengine => github.com/golang/appengine v1.3.0

 google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20181219182458-5a97ab628bfb

 google.golang.org/grpc => github.com/grpc/grpc-go v1.17.0

 gopkg.in/alecthomas/kingpin.v2 => github.com/alecthomas/kingpin v2.2.6+incompatible

 gopkg.in/mgo.v2 => github.com/go-mgo/mgo v0.0.0-20180705113604-9856a29383ce

 gopkg.in/vmihailenco/msgpack.v2 => github.com/vmihailenco/msgpack v2.9.1+incompatible

 gopkg.in/yaml.v2 => github.com/go-yaml/yaml v0.0.0-20181115110504-51d6538a90f8

 labix.org/v2/mgo => github.com/go-mgo/mgo v0.0.0-20160801194620-b6121c6199b7

 launchpad.net/gocheck => github.com/go-check/check v0.0.0-20180628173108-788fd7840127
```

### mongodb bson类型：
D: A BSON document. This type should be used in situations where order matters, such as MongoDB commands.
D是BSON文档的有序表示。当元素的顺序重要时，应使用此类型，
例如MongoDB命令文档。如果元素的顺序无关紧要，则应使用M代替。
用法示例：
bson.D {{"foo","bar"},{"hello","world"},{"pi",3.14159}}

M: An unordered map. It is the same as D, except it does not preserve order.
M是BSON文档的无序表示。当元素的顺序不正确时,应使用此类型
很重要。在编码和解码时，将这种类型作为常规map [string] interface {}处理。元素将是
以未定义的随机顺序序列化。如果元素的顺序很重要，则应使用D代替。
用法示例：
bson.M {"foo"："bar","hello"："world","pi"：3.14159}

A: A BSON array.A是BSON数组的有序表示。
用法示例：
bson.A {"bar","world",3.14159,bson.D {{"qux",12345}}}

E: A single element inside a D.
E代表D的BSON元素。通常在D内使用。

### 连接atlas数据库
```
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
client, err := mongo.Connect(ctx, options.Client().ApplyURI(
	"mongodb+srv://<username>:<password>@<host>/<database>?retryWrites=true&w=majority",
))
if err != nil { log.Fatal(err) }
err=client.Ping(context.TODO(), readpref.Primary())
if err!=nil {
	log.Fatal(err)
}
log.Println("Connected to MongoDB!")

u:=user.User{}
client.Database("website").Collection("website").FindOne(ctx, bson.M{}).Decode(u)
log.Println(u)
```