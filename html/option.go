// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package html

import (
	"myitcv.io/react"
)

// OptionElem is the React element definition corresponding to the HTML <option> element
type OptionElem struct {
	react.Element
}

// _OptionProps defines the properties for the <option> element
type _OptionProps struct {
	*BasicHTMLElement

	Value string `js:"value"`
}

// Option creates a new instance of a <option> element with the provided props
func Option(props *OptionProps, child react.Element) *OptionElem {

	rProps := &_OptionProps{
		BasicHTMLElement: newBasicHTMLElement(),
	}

	if props != nil {
		props.assign(rProps)
	}

	return &OptionElem{
		Element: react.InternalCreateElement("option", rProps, child),
	}
}
