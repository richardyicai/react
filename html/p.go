// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package html

import (
	"myitcv.io/react"
)

// PElem is the React element definition corresponding to the HTML <p> element
type PElem struct {
	react.Element
}

// _PProps are the props for a <div> component
type _PProps struct {
	*BasicHTMLElement
}

// P creates a new instance of a <p> element with the provided props and
// children
func P(props *PProps, children ...react.Element) *PElem {

	rProps := &_PProps{
		BasicHTMLElement: newBasicHTMLElement(),
	}

	if props != nil {
		props.assign(rProps)
	}

	return &PElem{
		Element: react.InternalCreateElement("p", rProps, children...),
	}
}
