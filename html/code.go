// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package html

import (
	"myitcv.io/react"
)

// CodeElem is the React element definition corresponding to the HTML <code> element
type CodeElem struct {
	react.Element
}

// _CodeProps defines the properties for the <code> element
type _CodeProps struct {
	*BasicHTMLElement
}

// Code creates a new instance of a <code> element with the provided props
func Code(props *CodeProps, children ...react.Element) *CodeElem {

	rProps := &_CodeProps{
		BasicHTMLElement: newBasicHTMLElement(),
	}

	if props != nil {
		props.assign(rProps)
	}

	return &CodeElem{
		Element: react.InternalCreateElement("code", rProps, children...),
	}
}
