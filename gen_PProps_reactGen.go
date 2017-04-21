// Code generated by reactGen. DO NOT EDIT.

package react

// PProps are the props for a <div> component
type PProps struct {
	ClassName               string
	DangerouslySetInnerHTML *DangerousInnerHTMLDef
	ID                      string
	Key                     string
	OnChange                func(e *SyntheticEvent)
	OnClick                 func(e *SyntheticMouseEvent)
	Role                    string
	Style                   *CSS
}

func (p *PProps) assign(v *_PProps) {

	v.ClassName = p.ClassName

	v.DangerouslySetInnerHTML = p.DangerouslySetInnerHTML

	if p.ID != "" {
		v.ID = p.ID
	}

	if p.Key != "" {
		v.Key = p.Key
	}

	v.OnChange = p.OnChange

	v.OnClick = p.OnClick

	v.Role = p.Role

	// TODO: until we have a resolution on
	// https://github.com/gopherjs/gopherjs/issues/236
	v.Style = p.Style.hack()

}