// Code generated by qtc from "transcript2wbpage.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line transcript2wbpage.qtpl:1
package diarization

//line transcript2wbpage.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line transcript2wbpage.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line transcript2wbpage.qtpl:1
func StreamTranscriptWebpage(qw422016 *qt422016.Writer, txn *Transcript) {
//line transcript2wbpage.qtpl:1
	qw422016.N().S(`<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<link href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css" rel="stylesheet" crossorigin="anonymous">
</head>
<body>
<h1>Transcript</h1>

<div style="margin:0em 5em 0em 5em">

`)
//line transcript2wbpage.qtpl:12
	qw422016.N().S(TranscriptHtml(txn))
//line transcript2wbpage.qtpl:12
	qw422016.N().S(`

</div>

</body>
</html>
`)
//line transcript2wbpage.qtpl:18
}

//line transcript2wbpage.qtpl:18
func WriteTranscriptWebpage(qq422016 qtio422016.Writer, txn *Transcript) {
//line transcript2wbpage.qtpl:18
	qw422016 := qt422016.AcquireWriter(qq422016)
//line transcript2wbpage.qtpl:18
	StreamTranscriptWebpage(qw422016, txn)
//line transcript2wbpage.qtpl:18
	qt422016.ReleaseWriter(qw422016)
//line transcript2wbpage.qtpl:18
}

//line transcript2wbpage.qtpl:18
func TranscriptWebpage(txn *Transcript) string {
//line transcript2wbpage.qtpl:18
	qb422016 := qt422016.AcquireByteBuffer()
//line transcript2wbpage.qtpl:18
	WriteTranscriptWebpage(qb422016, txn)
//line transcript2wbpage.qtpl:18
	qs422016 := string(qb422016.B)
//line transcript2wbpage.qtpl:18
	qt422016.ReleaseByteBuffer(qb422016)
//line transcript2wbpage.qtpl:18
	return qs422016
//line transcript2wbpage.qtpl:18
}
