+++
title = "30 seconds of Reactive code"
+++


借鉴 30-seconds-of-code 系列，将Reactive使用到的代码场景整理一下，方便大家参考。


# Reactive基础

### Publisher是Immutable(不变的)
Mono/Flux这些变量都Immutable的，也就是你每对其进行一次操作，会生成一个新的变量，而不是这期的变量，这个和Java中通常的处理不一样，如以下代码：

```java
public void updateInfo(User user) {
    user.nick = "leijuan";
    ...
}
```

上述代码中，user对象属性会被改变的，但是如果你将一个Flux对象传入一个void函数，那么是不会改变Flux的，如果你有这个需求，你需要使用返回值。如下：

```java
//add auditioin
public Flux<User> addAudition(Flux<User> flux) {
    return flux.doOnNext()...
}

//use transform
yourFlux.transform(flux->addAudition(flux))
                .map(
```


# Time时间

#### 全局定时器
之前你需要写一个Timer，然后做定时，现在只需要subscriber一个interval的Flux就可以啦。

```java
    @Bean
    Flux<Long> fiveSecondsTimer() {
        return Flux.interval(Duration.ofSeconds(3));
    }
```

#### 延迟消息消费
在做最终一致性检查时候作用比较大。 如收到买家购买旺铺旗舰版的服务，资金平台调用店铺API，但是不能确认店铺的是否为调整到新的状态啦，你可以设置一个延迟消息，调用店铺API进行检查。 使用delaySequence就可以，表示消息进入processor到被消费，进行一段时间的延迟。


```java
        DirectProcessor<Long> processor = DirectProcessor.create();
        processor.delaySequence(Duration.ofSeconds(15)).subscribe(t -> {
             //检查店铺状态
        });
        //设置要检查的店铺ID
        processor.onNext(1114L);

```

#### 消息间消费延迟
如果你担心消息消费的太快，细水长流消费，可以设置一个消息间消费延迟，如下述代码是100毫秒延迟。

```java
        Flux<String> flux = Flux.just("red", "White", "blue").delayElements(Duration.ofMillis(100));
```
同时，你可以设置delayUntil()，这样某些条件触发后才能进行消费。如你要获取某一隐私数据，只有安全接口审核通过后你才能获取到； 或者只有到某一个时间点大家才能收到消息，如双11的12点前5秒你才能得到消息。


#### 业务操作超时设置
普通的超时非常简单，设置一下timeout就可以啦，如下：

```java
Mono.just(1).timeout(Duration.ofSeconds(1000))
```

还有一些场景是发出去，然后等着回来，如果在等的内没有得到返回，则进行超时设置，下面代码是针对没有实现Reactive化的API的。 在create的逻辑中进行一些相关的操作，当如是CPU密集型也没有关系，如某一算法等，如果在规定的时间没有完成，直接timeout。

```java
   MonoProcessor.create((sink -> {
           //逻辑操作，然后将结果返回给sink
          //sink.success(1);
        })).timeout(Duration.ofSeconds(2)).doOnError((ex)->{
            System.out.println(ex.getMessage());
        }).subscribe();
```
如在做Node调用WebAssembly函数时候，你就可以使用这个方法，这样可以保证不会出现长时间等待空耗时间的问题。


# 钩子场景

### 清空Cache
如更新用户信息后清空Cache或者更新值，都可以用doOnNext()进行相应的更新，可以实现更多自定义逻辑。

```java
 Mono.just(user).doOnNext(temp->{
            //基于user id清空cache
           //基于user email 清空cache
        }) ;
```

### 计数器
在登录成功后，可以调用钩子进行在线用户数或者连接的统计。

### cleanup
如你访问一个InputStream的Mono，可以通过doOnTerminate进行关闭。

```java
 Mono.just(inputStream).doOnTerminate(() ->{
            try{
                inputStream.close();
            }  catch (Exception e) {

            }
        });
```

### close通知
如果大家都想监听某一资源的关闭通知，那么创建一个MonoProcessor，然后大家都subscribe就可以啦，资源在关闭的时候，会调用该MonoProcessor的onComplete()进行关闭通知。

```java
MonoProcessor<Void> onClose = MonoProcessor.create();
// close operatioin
onClose.onComplete();
```

### Pool设计
Pool主要是borrow和return对应的资源，所以在使用完资源后，在doOnTerminate的钩子中进行资源归还操作。 Reactor Pool就是这个设计机制 https://github.com/reactor/reactor-pool

```java
pool.withPoolable(resource ->
             resource
                .createStatement().flatMapMany(st -> st.query("SELECT * FROM foo"))
               .map(row -> rowToJson(row))
).map(json -> sanitize(json));
```

### 审计或者对账
在一些场景中，如调用外部手机充值接口，你这个时候需要将调用状态记录下来，如调用外部合作上的API完成手机号码充值，在调用完成后，添加一个doOnSuccess的钩子进行记录，什么时候我调用你充值接口啦，给我返回交易ID等，如果出现不成功，可以用这条记录进行对账。  当然这个充值的场景中，我们会添加另外一个1分钟的延迟对账消息，使用交易ID调用供应商接口查询是否真的充值成功啦，然后更新系统的状态。

```java
 mono.doOnSuccess(text->{

        })
```


# Processor：数据处理

### Back Pressure
如果你订阅的flux会给你实时推送时非常多消息，使用limitRate(100)可以保证flux每次给你推送100条消息，消息消费完备后会再给你发送100条，不会将订阅方打垮。

```java
           flux.limitRate(100)
                  .subscribe(payload -> {
                      System.out.println(payload.getDataUtf8());
                  });
```

### 模拟Topic
如果你想模仿传统的Message Broker做一个Topic，可以发送和订阅消息，那么使用EmitterProcessor就可以啦。

```java
val emitter = EmitterProcessor.create<Int>()
        emitter
            .map { it + 1 }
            .subscribe { println(it) }
        emitter.onNext(1)
        emitter.onNext(2)
```

### 应用的配置项
应用配置项需要保存最新一次配置SNAPSHOT，所以通过ReplayProcessor.cacheLast()可以缓存最后一次推送的配置，这样所有新上来的应用也可以收到最后一次配置推送。

```java
 ReplayProcessor<String> config = ReplayProcessor.cacheLast();
```

### 集群拓扑结构更新
在Config推送的基础上，将单个String对象调整为Collection<String>，代表集群中所有服务器的地址列表。

```java
// 创建一个要推送集群数据的processor
ReplayProcessor<Collection<String>> urisProcessor = ReplayProcessor.cacheLast();

//其他对象会在构造函数中订阅该processor，来响应集群的变化
public LoadBalancedRSocket( Flux<Collection<String>> urisProcessor) {
        this.urisFactory.subscribe(this::refreshRsockets);
    }
```

### Map转换数据累加
如你访问一个卖家的店铺详情， 第一步更加店铺id找店铺，然后根据店铺的卖家id找卖家，然后更加卖家中的账号id找会员，然后根据会员id找到对应的头像等其他信息。 在这个过程中，你不需要创建大量的Java Bean，借助Tuple，就可以将这些对象都保留下来，然后提供给页面进行渲染。

```java
Mono.just(1)
                .map(id-> Tuples.of(id, 111))
                .map(tuple2-> Tuples.of(tuple2.getT1(),tuple2.getT2(), 111))
                .map(tuple3->{ });
```

### CSV的数据处理
这个场景主要是指一个Flux流中，第一个数据是原数据，而不是我们要处理的数据，但是原始数据也非常有用，所以我们要处理第一个元数据，然后是接下来的真实数据。 如CSV流，第一个数据是CVS column names，接下来是数据。 使用switchOnFirst进行转换。

```java
  Flux.just("id,name", "1,leijuan", "2,juven").switchOnFirst((signal, stringFlux) -> {
            System.out.println("First: " + signal.get());
            return stringFlux.skip(1);
        }).subscribe(text -> {
            System.out.println(text);
        });

````

### 时间段的buffer
Flux的buffer可以做时间段的buffer，也就是将时间段内的数据形成list。 举一个例子，用户登录后，我们会将登录后的用户形成一个User Flux，但是我有一个统计，指向统计一分钟内有多少用户登录，那么调用buffer就可以，然后将该时间段的user list的size记录并统计一下。

```java
loginUsers.buffer(Duration.ofMinutes(1))
```

### 基于buffer的小batch
就是将流式的数据形成小batch后处理。 举一个例子，我们答应商家和运营同学可以每次导出500个订单的CSV文件，但是我们调用交易中心的时候，每次只允许我们取20条记录，这个时候，500条记录就需要划分为25个batch，然后并发向交易中心查询，然后将结果进行合并，返回给卖家或者运营同学。

```java
 flux.buffer(20).map(idList -> {
           //调用接口，返回对象list
        }).
```

# Reactive Exception

Reactive中的异常处理和我们通常理解的try-catch有一定的区别，事实上更方便理解。

### 异常捕获
在Reactor框架中主要有四个操作符来处理异常doOnError, onErrorMap, onErrorReturn,和onErrorResume

* doOnError: 当异常发生时会执行该操作，但是异常不会被捕获，还是会继续抛给最终消费方。主要的场景如做错误日志记录等。
* onErrorMap: 将发生的异常转换为另外一个异常，然后抛出转换后的异常。主要场景如将IO异常转换为业务异常，更方便消费方理解或者网络传输。
* onErrorReturn: 当异常发生时，会返回指定的缺省值，异常会被捕获，不会继续抛出。 主要的场景是用缺省值方式来替换异常抛出。
* onErrorResume: 当异常发生时，会调用fallback Reactive函数，然后将函数返回的值以flatMap方式返回给调用方。


### 异常抛出
前面讲到如何捕获异常，那么在实际的代码中如何抛出异常？ 传统的throw方式在Reactive中要被抛弃，如以下代码千万不要使用：

```java
 Mono.just("https://www.taobao.com/").map(text -> {
            try {
                return new URI(text);
            } catch (Exception e) {
                throw new RuntimeException(e);
            }
        })
```

接下来我们就介绍一下常见异常抛出的方式：

* handle:  handle函数提供一个sink，我们可以直接调用sink.error()，可以处理各种复杂异常。

```java
 Mono.just("https://www.taobao.com/").handle((text, sink) -> {
            try {
                sink.next(new URI(text));
            } catch (Exception e) {
                sink.error(e);
            }
        })
```

* concatMap + Mono.error: 处理转换中的异常，如String转换为URI对象等。

```java
Mono.just(userId)
    .map(repo::findById)
    .concatMap(user-> {
        if(!isValid(user)){
            return Mono.error(new InvalidUserException());
        }
        return Mono.just(user);
    })
```

当然flatMap也可也做同样的事情，但是这个场景下contactMap更适合，contact是转换操作，而flatMap是做多个流式合并。

* switchOnEmpty: 不少情况下我们希望在空值的情况下抛出异常，如典型的NotFoundException

```
Mono.just(userId)
    .flatMap(repo::findById)
    .switchIfEmpty(Mono.error(new UserNotFoundExeception()))

```

* flatMap + Mono.error: 合并流操作的时候，可以抛出异常。 注意这里是合并多个流。

当然Reactive中是不允许空值的，如果流中包含null值，会直接抛出 NullPointerException，这个你可能要进行处理。 如果你确认值可能会Null，请调用  Mono.justOrEmpty()

### 空值(empty)处理
虽然Empty和Exception不太一样，这里还是放在一起方便理解。当我们遇到Reactive中empty时，会有一些方法来方便我们处理：

* defaultIfEmpty: 非常容易理解，如果为空我们使用一个缺省值代替
* switchIfEmpty: 使用另外一个Mono或者Flux来代替
* repeatWhenEmpty: 如果为空，则重复执行再次订阅，直到有非空值返回。 如下面代码，如果为空值，则再次发起订阅，那么map, flatMap都会被重新执行3次(最大重复数是5次)，直到第四次返回非空值。

```java
 AtomicInteger atomicInteger = new AtomicInteger(1);
 Mono.just(0)
      .map(num -> {
          System.out.println("map: " + atomicInteger.get());
          return num;
      })
      .flatMap(text -> {
          System.out.println("flatMap: ");
          if (atomicInteger.incrementAndGet() <= 3) {
              return Mono.empty();
          }
          return Mono.just(atomicInteger.get());
      })
      .repeatWhenEmpty(Repeat.times(5));
```

# Reactor Context

### Mutable Context 可变的Context
Context这个是Reactor支持的一个特性，也就是注入上下文，可以绑定到一个Reactive的执行链上，但是默认是只读的，如果你有一些需要，要做各个filter，flatmap做一些数据调整，可以考虑使用MutableContext，代码如下：

```java
public class MutableContext implements Context {
    HashMap<Object, Object> holder = new HashMap<>();

    @SuppressWarnings("unchecked")
    @NotNull
    @Override
    public <T> T get(@NotNull Object key) {
        return (T) holder.get(key);
    }

    @Override
    public boolean hasKey(@NotNull Object key) {
        return holder.containsKey(key);
    }

    @NotNull
    @Override
    public Context put(@NotNull Object key, @NotNull Object value) {
        holder.put(key, value);
        return this;
    }

    @NotNull
    @Override
    public Context delete(@NotNull Object key) {
        holder.remove(key);
        return this;
    }

    @Override
    public int size() {
        return holder.size();
    }

    @NotNull
    @Override
    public Stream<Map.Entry<Object, Object>> stream() {
        return holder.entrySet().stream();
    }

}
```

样例代码如下：

```java
  Mono.just("Hello")
                .flatMap(s -> Mono.subscriberContext()
                        .map(ctx -> {
                            return s + " " + ctx.get("nick");
                        }))
                .subscriberContext(ctx -> ctx.put("nick", "Reactor"))
                .subscriberContext(context)
```

### 执行前先从Context中获取变量
一些情况下，我们要从Context获取特定变量信息，然后进行逻辑执行，如获取当前访问用户的nick，这个时候使用deferWithContext即可。

```java
 Mono.deferWithContext(ctx -> Mono.just(ctx.get("nick")))
```

### Threadlocal的问题
如果你确实有一些Threadlocal的代码，确实要需要使用它，那么可以将thread local变量转换为Context进行处理：

```java
 MutableContext context = new MutableContext();
        context.put("nick", userThreadLocal.get());
        Mono.deferWithContext(ctx -> Mono.just(ctx.get("nick")))
                .subscriberContext(context)
```

# 其他

### 空值处理
如果API返回为Mono，如 Mono<User>，则表示可能会出现空值的情况，也就是返回 Mono.empty()，这个时候，如果你想使用缺省值(default value)，可以调用then或者defaultIfEmpty(Object)

```java
Mono.empty().then(Mono.just("default"))
```

### 标签支持(Tag support)
有些时候我们在返回一个标准的对象，如Mono<User>，我们还希望附加一些标签支持，这个时候你可以通过tag添加字符串标签，代码如下：

```java
//调用tag进行打标
Mono<String> nick = Mono.just("leijuan").tag("alias", "linux_china");
//调用Scannable获取标签
Map<String, String> tags = Scannable.from(nick).tags().collect(Collectors.toMap(Tuple2::getT1, Tuple2::getT1));
```

这个不是标准的Reactive特性，请酌情使用。  另外还有一个name()的方法，可以让你取得publisher的name，你可以进行相应的逻辑判断。

### Cache支持
如果想将Mono/Flux的值作为Cache缓存起来，然后提供给其他消费方进行消费，那么调用cache()就可以，然后Data, Completion and Error都会被重放。如果你想设置全局Mono对象或者Cache支持，这个方法不错。

```java
 Mono<String> user = Mono.<String>create(monoSink -> {
            System.out.println("Only Once");
            monoSink.success("nick");
        }).cache();
        user.subscribe(t -> {
            System.out.println(t);
        });
        user.subscribe(t -> {
            System.out.println(t);
        });
```
当然cache还提供ttl支持，如果你想设置ttl，也没有问题，这样你可以将远程或者数据库返回的调用结果进行缓存。

### 完全Lazy
如果调用一个函数，该函数的返回值是Mono，但是在函数调用的过程中，还是会执行函数中的同步代码，如以下代码，System.out.println还是会被执行的。

```java
public Mono<String> getNick() {
        System.out.println("you are invoking getNick()");
        return Mono.just("nick");
}

```

如果你想完全是lazy的，等待subscribe的时候再执行该函数，那么使用defer就可以，代码如下：

```java
Mono<String> defer = Mono.defer(this::getNick);
```

### Reactive框架之间的互操作
Reactor Adapter可以让RxJava, Akka, CompletableFuture之间都是相互转换的，即便之前使用RxJava或者CompletableFuture，都是可以和Reactor互操作的，而且Reactor也能转换为RxJava接口。

### 参考

* ReactiveX Operators: http://reactivex.io/documentation/operators.html
* Project Reactor Operators: https://projectreactor.io/docs/core/release/reference/#which-operator
* Learn RxJS: https://www.learnrxjs.io/ https://rxjs-cn.github.io/learn-rxjs-operators/
* Interactive diagrams of Rx Observables: https://rxmarbles.com/
* RxJava Operator Matrix: https://github.com/ReactiveX/RxJava/wiki/Operator-Matrix
