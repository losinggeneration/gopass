package main

// This file is automatically generated by gopkg.in/qml.v1/cmd/genqrc

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/qml.v1"
)

func init() {
	var r *qml.Resources
	var err error
	if os.Getenv("QRC_REPACK") == "1" {
		err = qrcRepackResources()
		if err != nil {
			panic("cannot repack qrc resources: " + err.Error())
		}
		r, err = qml.ParseResources(qrcResourcesRepacked)
	} else {
		r, err = qml.ParseResourcesString(qrcResourcesData)
	}
	if err != nil {
		panic("cannot parse bundled resources data: " + err.Error())
	}
	qml.LoadResources(r)
}

func qrcRepackResources() error {
	subdirs := []string{"assets/main.qml", "assets/logo.svg"}
	var rp qml.ResourcesPacker
	for _, subdir := range subdirs {
		err := filepath.Walk(subdir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}
			data, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			rp.Add(filepath.ToSlash(path), data)
			return nil
		})
		if err != nil {
			return err
		}
	}
	qrcResourcesRepacked = rp.Pack().Bytes()
	return nil
}

var qrcResourcesRepacked []byte
var qrcResourcesData = "qres\x00\x00\x00\x01\x00\x00'\x9c\x00\x00\x00\x14\x00\x00'^\x00\x00\x16\xf4<?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"no\"?>\n<!-- Created with Inkscape (http://www.inkscape.org/) -->\n\n<svg\n   xmlns:dc=\"http://purl.org/dc/elements/1.1/\"\n   xmlns:cc=\"http://creativecommons.org/ns#\"\n   xmlns:rdf=\"http://www.w3.org/1999/02/22-rdf-syntax-ns#\"\n   xmlns:svg=\"http://www.w3.org/2000/svg\"\n   xmlns=\"http://www.w3.org/2000/svg\"\n   xmlns:xlink=\"http://www.w3.org/1999/xlink\"\n   xmlns:sodipodi=\"http://sodipodi.sourceforge.net/DTD/sodipodi-0.dtd\"\n   xmlns:inkscape=\"http://www.inkscape.org/namespaces/inkscape\"\n   width=\"64\"\n   height=\"64\"\n   id=\"svg4265\"\n   version=\"1.1\"\n   inkscape:version=\"0.91 r13725\"\n   viewBox=\"0 0 64 64\"\n   sodipodi:docname=\"logo.svg\">\n  <defs\n     id=\"defs4267\">\n    <linearGradient\n       inkscape:collect=\"always\"\n       id=\"linearGradient4891\">\n      <stop\n         style=\"stop-color:#e6e6e6;stop-opacity:1\"\n         offset=\"0\"\n         id=\"stop4893\" />\n      <stop\n         style=\"stop-color:#575757;stop-opacity:1\"\n         offset=\"1\"\n         id=\"stop4895\" />\n    </linearGradient>\n    <radialGradient\n       inkscape:collect=\"always\"\n       xlink:href=\"#linearGradient4891\"\n       id=\"radialGradient4901\"\n       cx=\"28.997917\"\n       cy=\"10.516196\"\n       fx=\"28.997917\"\n       fy=\"10.516196\"\n       r=\"18.54099\"\n       gradientTransform=\"matrix(0.79739441,-0.66023391,0.48669645,0.58780536,-1.8659729,26.464997)\"\n       gradientUnits=\"userSpaceOnUse\" />\n  </defs>\n  <sodipodi:namedview\n     id=\"base\"\n     pagecolor=\"#ffffff\"\n     bordercolor=\"#666666\"\n     borderopacity=\"1.0\"\n     inkscape:pageopacity=\"0.0\"\n     inkscape:pageshadow=\"2\"\n     inkscape:zoom=\"8\"\n     inkscape:cx=\"-5.6325773\"\n     inkscape:cy=\"32.144434\"\n     inkscape:current-layer=\"layer1\"\n     showgrid=\"true\"\n     inkscape:document-units=\"px\"\n     inkscape:grid-bbox=\"true\"\n     showguides=\"false\"\n     inkscape:guide-bbox=\"true\"\n     inkscape:window-width=\"2560\"\n     inkscape:window-height=\"1376\"\n     inkscape:window-x=\"0\"\n     inkscape:window-y=\"27\"\n     inkscape:window-maximized=\"1\" />\n  <metadata\n     id=\"metadata4270\">\n    <rdf:RDF>\n      <cc:Work\n         rdf:about=\"\">\n        <dc:format>image/svg+xml</dc:format>\n        <dc:type\n           rdf:resource=\"http://purl.org/dc/dcmitype/StillImage\" />\n        <dc:title></dc:title>\n      </cc:Work>\n    </rdf:RDF>\n  </metadata>\n  <g\n     inkscape:groupmode=\"layer\"\n     id=\"layer3\"\n     inkscape:label=\"Hanger\"\n     style=\"display:inline\">\n    <path\n       style=\"display:inline;opacity:1;fill:none;fill-opacity:1;stroke:url(#radialGradient4901);stroke-width:5;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1\"\n       d=\"m 32.5,4.731595 c 10.648,0 16.620529,5.381024 16,12 l -1.5,16 c -1.949438,17.87752 -28.552416,14.821066 -29,0 l -1.5,-16 c -0.620529,-6.618976 5.352,-12 16,-12 z\"\n       id=\"rect4871\"\n       inkscape:connector-curvature=\"0\"\n       sodipodi:nodetypes=\"csccsc\" />\n  </g>\n  <g\n     id=\"layer1\"\n     inkscape:label=\"Layer 1\"\n     inkscape:groupmode=\"layer\"\n     style=\"display:inline\">\n    <rect\n       style=\"fill:#ffaaee;fill-opacity:1;stroke:none;stroke-width:1;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1\"\n       id=\"rect4320\"\n       width=\"13.370747\"\n       height=\"11.827968\"\n       x=\"77.524612\"\n       y=\"2.1602979\" />\n    <rect\n       y=\"16.173868\"\n       x=\"77.396049\"\n       height=\"11.827968\"\n       width=\"13.370747\"\n       id=\"rect4322\"\n       style=\"fill:#c6afe9;fill-opacity:1;stroke:none;stroke-width:1;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1\" />\n    <rect\n       style=\"fill:#8d5fd3;fill-opacity:1;stroke:none;stroke-width:1;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1\"\n       id=\"rect4324\"\n       width=\"13.370747\"\n       height=\"11.827968\"\n       x=\"77.396049\"\n       y=\"30.701698\" />\n    <g\n       id=\"g4932\"\n       clip-path=\"none\">\n      <rect\n         y=\"24\"\n         x=\"8\"\n         height=\"34\"\n         width=\"50\"\n         id=\"rect4318\"\n         style=\"fill:#8d5fd3;fill-opacity:1;stroke:none;stroke-width:1;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1\" />\n      <path\n         style=\"fill:#c6afe9;fill-opacity:1;stroke:none;stroke-width:1;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1\"\n         d=\"M 58,24 8,58 58,58 Z\"\n         id=\"path4329\"\n         inkscape:connector-curvature=\"0\"\n         sodipodi:nodetypes=\"cccc\" />\n      <path\n         sodipodi:nodetypes=\"cccc\"\n         inkscape:connector-curvature=\"0\"\n         id=\"rect4326\"\n         d=\"M 8,24 58,58 8,58 Z\"\n         style=\"opacity:0.75;fill:#ffaaee;fill-opacity:1;stroke:none;stroke-width:1;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1\"\n         clip-path=\"none\"\n         mask=\"none\" />\n    </g>\n  </g>\n  <g\n     inkscape:groupmode=\"layer\"\n     id=\"layer2\"\n     inkscape:label=\"Keyhole\"\n     style=\"display:inline\"\n     sodipodi:insensitive=\"true\">\n    <circle\n       style=\"display:inline;opacity:1;fill:#2e1f45;fill-opacity:1;stroke:none;stroke-width:5.28483486;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1\"\n       id=\"path4873\"\n       cx=\"32.59758\"\n       cy=\"38.147919\"\n       r=\"5.6146016\" />\n    <path\n       sodipodi:type=\"star\"\n       style=\"display:inline;opacity:1;fill:#2e1f45;fill-opacity:1;stroke:none;stroke-width:5;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1\"\n       id=\"path4877\"\n       sodipodi:sides=\"3\"\n       sodipodi:cx=\"34.32682\"\n       sodipodi:cy=\"57.571754\"\n       sodipodi:r1=\"11.673945\"\n       sodipodi:r2=\"5.8369727\"\n       sodipodi:arg1=\"0.52359878\"\n       sodipodi:arg2=\"1.5707963\"\n       inkscape:flatsided=\"false\"\n       inkscape:rounded=\"0\"\n       inkscape:randomized=\"0\"\n       d=\"m 44.436754,63.408727 -10.109933,0 -10.109934,0 5.054967,-8.755459 5.054966,-8.755459 5.054967,8.755459 z\"\n       inkscape:transform-center-y=\"-2.3907005\"\n       transform=\"matrix(0.6284863,0,0,0.81915906,11.179559,-2.2339567)\" />\n  </g>\n</svg>\n\x00\x00\x10Nimport QtQuick 2.5\nimport QtQuick.Controls 1.3\nimport QtQuick.Layouts 1.2\nimport QtQuick.Controls.Styles 1.4\nimport QtQuick.Window 2.2\n\nApplicationWindow {\n    id: rootWindow\n    visible: true\n    title: \"GoPass\"\n    property int margin: 10\n    width: 600\n    height: mainLayout.implicitHeight + 2 * margin\n    minimumWidth: mainLayout.Layout.minimumWidth + 2 * margin\n    minimumHeight: 400\n\n    x: (Screen.width - width) / 2\n\n    flags: Qt.FramelessWindowHint | Qt.Window\n    color: \"transparent\"\n\n    Rectangle {\n        color: \"#333\"\n        anchors.fill: parent\n        anchors.margins: 0\n        radius: 10\n        border.width: 2\n        border.color: \"#aaa\"\n    }\n\n    MouseArea {\n        id: mouseRegion\n        anchors.fill: parent;\n        property variant clickPos: \"1,1\"\n\n        onPressed: {\n            clickPos  = Qt.point(mouse.x,mouse.y)\n        }\n\n        onPositionChanged: {\n            var delta = Qt.point(mouse.x-clickPos.x, mouse.y-clickPos.y)\n            rootWindow.x += delta.x;\n            rootWindow.y += delta.y;\n        }\n    }\n\n    Shortcut {\n        sequence:\"Ctrl+K\"\n        context: Qt.ApplicationShortcut\n        onActivated: passwords.up()\n    }\n\n    Shortcut {\n        sequence:\"Up\"\n        context: Qt.ApplicationShortcut\n        onActivated: passwords.up()\n    }\n\n    Shortcut {\n        sequence:\"Ctrl+j\"\n        onActivated: passwords.down()\n    }\n\n    Shortcut {\n        sequence:\"Down\"\n        onActivated: passwords.down()\n    }\n\n    Shortcut {\n        sequence:\"Ctrl+r\"\n        onActivated: passwords.clearmetadata()\n    }\n\n    Shortcut {\n        sequence:\"Ctrl+l\"\n        onActivated: searchInput.selectAll()\n    }\n\n    Shortcut {\n        sequence: \"Esc\"\n        context: Qt.ApplicationShortcut\n        onActivated: passwords.quit()\n    }\n\n    Component {\n        id: passwordEntry\n        Text {\n            text: passwords.password(index).name;\n            font.pixelSize: 24\n            color: index==passwords.selected ? \"#dd00bb\": \"gray\"\n        }\n    }\n\n    ColumnLayout {\n        id: mainLayout\n        anchors.fill: parent\n        anchors.margins: margin\n        RowLayout{\n            Layout.fillWidth: true\n            TextField {\n                id: searchInput\n                font.pixelSize: 24\n                focus: true\n                onTextChanged: passwords.query(text)\n                onAccepted: passwords.copyToClipboard()\n                textColor: \"white\"\n                style: TextFieldStyle {\n                    textColor: \"white\"\n                    background: Rectangle {\n                        radius: 5\n                        border.color: \"#666\"\n                        border.width: 1\n                        color: \"#333\"\n                    }\n                }\n                Layout.fillWidth: true\n            }\n            Image {\n                source: \"logo.svg\"\n                fillMode: Image.PreserveAspectFit\n                Layout.alignment: Qt.AlignRight\n                Layout.maximumWidth: 32\n            }\n        }\n        ScrollView{\n            Layout.fillHeight: true\n            Layout.fillWidth: true\n            style: ScrollViewStyle{\n                transientScrollBars: true\n                scrollToClickedPosition : true\n            }\n            ListView {\n                id: hitList\n                model: passwords.len\n                delegate: passwordEntry\n                Layout.fillHeight: true\n            }\n        }\n        Text {\n            id: status\n            text: passwords.status\n            font.pixelSize: 14\n            color: \"#aaa\"\n        }\n\n        ScrollView{\n            Layout.fillHeight: true\n            Layout.fillWidth: true\n            Layout.preferredHeight: metadata.implicitHeight\n            style: ScrollViewStyle{\n                transientScrollBars: true\n                scrollToClickedPosition : true\n            }\n\n            TextEdit {\n                id: metadata\n                selectByMouse: true\n                readOnly: true\n                text: passwords.metadata\n                font.pixelSize: 14\n                color: \"#fa5858\"\n                wrapMode: TextEdit.WordWrap\n            }\n        }\n    }\n}\n\x00\x06\x06\x8a\x9c\xb3\x00a\x00s\x00s\x00e\x00t\x00s\x00\b\x05\xe2T\xa7\x00l\x00o\x00g\x00o\x00.\x00s\x00v\x00g\x00\b\b\x01Z\\\x00m\x00a\x00i\x00n\x00.\x00q\x00m\x00l\x00\x00\x00\x00\x00\x02\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x00\x00\x02\x00\x00\x00\x02\x00\x00\x00\x02\x00\x00\x00\x12\x00\x00\x00\x00\x00\x01\x00\x00\x00\x00\x00\x00\x00(\x00\x00\x00\x00\x00\x01\x00\x00\x16\xf8"
