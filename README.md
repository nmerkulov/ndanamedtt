# ndanamedtt


Sadly, but i ran out of time
Must say that i had very few experience with grpc itself (i had production experience with plenty of rest-ish and rest-full and even twirp apis) so most of 2 hours
i tried to remember how that works, from code perspective.

Also i didn't found quite a lot of behavior in this task so all of the domains contains "Anemic domain", which is anti pattern, according to mr. Fowler.

However i applied clean/hex architecture as a "sort-of framework" for DDD. You may found basic concepts implemented

I included golangci-linter, makefile, docker, protolock (for backward compatibility check) and utilized code generation by "go generate"

Code itself isn't complete. Things left to implement:
Inmemory storage for portDomainService (or sql implementation using sqlx + squirrel as most lightweight and easiest approach for small-medium size projects)
Wire everything up in main. I follow dependency injection approach without any service-locators or anything like that
Tracing. It is extremely important in microservice world
Authenticating between services
Metrics. Metrics are crucial in terms of using your software.
And service-discovery might be a thing

I decided to stop at 2 hours state as it is.

Thank you for your time,i'd appreciate if you ask some questions regarding my task
