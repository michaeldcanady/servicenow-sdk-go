package tableapi

type View2 int64

const (
    ViewDesktop View2 = iota
    ViewMobile
    ViewBoth
)

func (v View2) String() string {
    return map[View2]string{
        ViewDesktop: "desktop",
        ViewMobile: "mobile",
        ViewBoth: "both",
    }[v]
}
