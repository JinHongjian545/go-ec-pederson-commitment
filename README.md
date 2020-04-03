# go-ec-pederson-commitment
This is a Go implementation of the EC-based Pederson commitment scheme

golang 实现的EC pederson-commitment

EC 曲线使用的是 Edwards25519，由 ristretto 包提供，文档参考 [https://pkg.go.dev/github.com/bwesterb/go-ristretto?tab=doc ]() ，GitHub库 [ https://github.com/bwesterb/go-ristretto ](https://github.com/bwesterb/go-ristretto)

另外参考了 [ https://github.com/threehook/go-pedersen-commitment ](https://github.com/threehook/go-pedersen-commitment) 的实现，均已star

这里的封装是为了实现我自己的一些需求，同时也分享出来给需要的朋友

pederson-commitment 原论文可见 [ https://link.springer.com/content/pdf/10.1007%2F3-540-46766-1_9.pdf ](https://link.springer.com/content/pdf/10.1007%2F3-540-46766-1_9.pdf)