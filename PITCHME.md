# Boxes vs. Lego

![Box vs. Lego](assets/boxVsLego.png)

---
## Boxes are awesome

![Box](assets/box.png)

Note:
- They have a simple interface
- They hide details
- They can be used anywhere
- They can be put into other boxes

---
## Functions are like boxes

```go
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
- They can be used anywhere |
- They can be called from other functions |

---
## Boxes are awful

![Mess of boxes](assets/messyBoxes.jpg)

Note:
- Padding is need for nesting
- If you nest them deep enough it's hard to find something
- They can't be composed
- I have never seen a storage built out of boxes

---
## Functions are like boxes

```go
func TrimLeftFunc(s string, f func(rune) bool) string {
	i := indexFunc(s, f, false)
	if i == -1 {
		return ""
	}
	return s[i:]
}
```

- There are no standards |
- Glue code is needed for combining them |
- They can't be composed |
- I have seen too many programs built out of fuctions |

---
## Lego bricks are awesome

![Lego brick](assets/lego.png)

Note:
- They have a simple interface
- They have standardized sizes and forms
- They hide details
- They can be used almost anywhere
- They combine with other Lego bricks

---
## Graphical representation

![Simple operation](assets/simpleOp.png)

- The interface is simple |
- It has got a standardized form |
- The implementation is hidden |
- It can be used almost anywhere |

---
## They combine to flows

![Minimal flow](assets/simpleFlow.png)

The combined ops look like a single op

---
## The code for operation

```go
type StringFilter func(out func(string)) (in func(string))
```

- The interface is simple |
- It has got a standardized form |
- The implementation is hidden |
- It can be used almost anywhere |

---
## The code for operation

```go
type StringFilter func(out func(string)) (in func(string))

func TrimLeft(out func(string)) (in func(string)) {
	in = func(s string) {
		t := string.TrimLeft(s, " \t\r\n")
		out(t)
	}
	return
}
```

---
## The code for flow

![Minimal flow](assets/simpleFlow.png)

```go
func Trim(out func(string)) (in func(string)) {
	trIn := TrimRight(out)
	in = TrimLeft(trIn)
	return
}
```

---
## References:

- FlowDev project: [https://github.com/flowdev](https://github.com/flowdev)
- Slides online: [https://gitpitch.com/flowdev/present-go-boxes-lego](https://gitpitch.com/flowdev/present-go-boxes-lego)

## Attribution:

<div>Icons made by <a href="http://www.freepik.com" title="Freepik">Freepik</a> and <a href="https://www.flaticon.com/authors/coucou" title="Coucou">Coucou</a> from <a href="https://www.flaticon.com/" title="Flaticon">www.flaticon.com</a> is licensed by <a href="http://creativecommons.org/licenses/by/3.0/" title="Creative Commons BY 3.0" target="_blank">CC 3.0 BY</a></div>

<div>Photo made by <a href="https://www.flickr.com/people/skrewtape/" title="Screwtape">Screwtape</a> from <a href="https://www.flickr.com/photos/skrewtape/851672959" title="Flickr">www.flickr.com</a> is licensed by <a href="https://creativecommons.org/licenses/by-sa/2.0/" title="Creative Commons BY-SA 2.0" target="_blank">CC BY-SA 2.0</a></div>
