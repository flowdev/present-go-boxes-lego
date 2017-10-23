# Boxes Versus Lego Bricks in Go

![Box](assets/box.png)

versus

![Lego brick](assets/lego.png)

---
## Boxes are great

![Box](assets/box.png)

- They have a simple interface |
- They hide details |
- They can be used anywhere |
- They can be put into other boxes |

---
## Functions are like boxes

```
func TrimLeftFunc(s string, f func(rune) bool) string {
	i := indexFunc(s, f, false)
	if i == -1 {
		return ""
	}
	return s[i:]
}
```

- They have a simple interface |
- They hide details |
- They can be used everywhere |
- They can be put into other functions |

---
## Boxes are awful

- You usually need padding for cascading |
- If you cascade them mutiple times it's hard to find something |
- They can't be composed |
- I have never seen a storage built out of boxes |
- I have even less seen a house built out of boxes |

---
## Lego bricks are ...

- bli ![Lego brick](assets/lego.png)

---
## References:

- FlowDev project: [https://github.com/flowdev](https://github.com/flowdev)
- Slides online: [https://gitpitch.com/flowdev/present-go-boxes-lego](https://gitpitch.com/flowdev/present-go-boxes-lego)

## Attribution:

<div>Icons made by <a href="http://www.freepik.com" title="Freepik">Freepik</a> and <a href="https://www.flaticon.com/authors/coucou" title="Coucou">Coucou</a> from <a href="https://www.flaticon.com/" title="Flaticon">www.flaticon.com</a> is licensed by <a href="http://creativecommons.org/licenses/by/3.0/" title="Creative Commons BY 3.0" target="_blank">CC 3.0 BY</a></div>

---

## Procedural Programming

Note:
A call decouples from
- registers and
- variables and state in the calling procedure.

A procedure can be called from and reused anywhere.

This is still used a lot esp. for utility functions.

---

## The Goal

- Decoupling as a natural thing to happen |
- Good match for business logic           |
- Easy to understand                      |
- Natural graphical representation        |

