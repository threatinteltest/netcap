/*
 * NETCAP - Traffic Analysis Framework
 * Copyright (c) 2017 Philipp Mieden <dreadl0ck [at] protonmail [dot] ch>
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

package types

import (
	"encoding/hex"
	"strings"
)

func (i ICMPv6RouterSolicitation) CSVHeader() []string {
	return filter([]string{
		"Timestamp",
		"Options",
	})
}

func (i ICMPv6RouterSolicitation) CSVRecord() []string {
	var opts []string
	for _, o := range i.Options {
		opts = append(opts, o.ToString())
	}
	return filter([]string{
		formatTimestamp(i.Timestamp),
		strings.Join(opts, ""),
	})
}

func (i ICMPv6RouterSolicitation) NetcapTimestamp() string {
	return i.Timestamp
}

func (o ICMPv6Option) ToString() string {

	var b strings.Builder
	b.WriteString(begin)
	b.WriteString(formatInt32(o.Type))
	b.WriteString(sep)
	b.WriteString(hex.EncodeToString(o.Data))
	b.WriteString(end)

	return b.String()
}
