package gemail

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/smtp"
	"path/filepath"
	"strings"
	"time"
)

// attachment ...
type attachment struct {
	Filename string
	Data     []byte
	Inline   bool
}

// Message ...
type Message struct {
	From string
	To   []string
	Cc   []string
	Bcc  []string

	subject         string
	body            string
	bodyContentType string
	attachments     map[string]*attachment
}

func (m *Message) attach(file string, inline bool) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	_, filename := filepath.Split(file)

	m.attachments[filename] = &attachment{
		Filename: filename,
		Data:     data,
		Inline:   inline,
	}

	return nil
}

// Attach a file to email by ouline
func (m *Message) Attach(file string) error {
	return m.attach(file, false)
}

// Inline attach a file to email by inline
func (m *Message) Inline(file string) error {
	return m.attach(file, true)
}

func newMessage(subject string, body string, bodyContentType string) *Message {
	m := &Message{subject: subject, body: body, bodyContentType: bodyContentType}
	m.attachments = make(map[string]*attachment)
	return m
}

// NewMessage returns a new Message that can compose an email with attachments
func NewMessage(subject string, body string) *Message {
	return newMessage(subject, body, "text/plain")
}

// NewHTMLMessage returns a new Message that can compose an HTML email with attachments
func NewHTMLMessage(subject string, body string) *Message {
	return newMessage(subject, body, "text/html")
}

// Tolist returns all the recipients of the email
func (m *Message) Tolist() []string {
	tolist := m.To

	for _, cc := range m.Cc {
		tolist = append(tolist, cc)
	}

	for _, bcc := range m.Bcc {
		tolist = append(tolist, bcc)
	}

	return tolist
}

// Bytes returns the mail data
func (m *Message) bytes() []byte {
	buf := bytes.NewBuffer(nil)

	buf.WriteString("From: " + m.From + "\n")

	t := time.Now()
	buf.WriteString("Date: " + t.Format(time.RFC822) + "\n")

	buf.WriteString("To: " + strings.Join(m.To, ",") + "\n")
	if len(m.Cc) > 0 {
		buf.WriteString("Cc: " + strings.Join(m.Cc, ",") + "\n")
	}

	buf.WriteString("Subject: " + m.subject + "\n")
	buf.WriteString("MIME-Version: 1.0\n")

	boundary := "f46d043c813270fc6b04c2d223da"

	if len(m.attachments) > 0 {
		buf.WriteString("Content-Type: multipart/mixed; boundary=" + boundary + "\n")
		buf.WriteString("--" + boundary + "\n")
	}

	buf.WriteString(fmt.Sprintf("Content-Type: %s; charset=utf-8\n\n", m.bodyContentType))
	buf.WriteString(m.body)
	buf.WriteString("\n")

	if len(m.attachments) > 0 {
		for _, attachment := range m.attachments {
			buf.WriteString("\n\n--" + boundary + "\n")

			if attachment.Inline {
				buf.WriteString("Content-Type: message/rfc822\n")
				buf.WriteString("Content-Disposition: inline; filename=\"" + attachment.Filename + "\"\n\n")

				buf.Write(attachment.Data)
			} else {
				buf.WriteString("Content-Type: application/octet-stream\n")
				buf.WriteString("Content-Transfer-Encoding: base64\n")
				buf.WriteString("Content-Disposition: attachment; filename=\"" + attachment.Filename + "\"\n\n")

				b := make([]byte, base64.StdEncoding.EncodedLen(len(attachment.Data)))
				base64.StdEncoding.Encode(b, attachment.Data)
				buf.Write(b)
			}

			buf.WriteString("\n--" + boundary)
		}

		buf.WriteString("--")
	}

	return buf.Bytes()
}

// Send message to smtp server
func (m *Message) Send(addr string, auth smtp.Auth) error {
	return smtp.SendMail(addr, auth, m.From, m.Tolist(), m.bytes())
}
