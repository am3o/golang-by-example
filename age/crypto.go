package main

import (
	"fmt"
	"io"

	"filippo.io/age"
)

type Container struct {
	secret string
	file   io.ReadWriter
}

func NewContainer(file io.ReadWriter, passphrase string) (Container, error) {
	return Container{
		secret: passphrase,
		file:   file,
	}, nil
}

func (c *Container) Reader() (io.Reader, error) {
	identity, err := age.NewScryptIdentity(c.secret)
	if err != nil {
		return nil, fmt.Errorf("could not create crypt identity: %w", err)
	}

	decrypt, err := age.Decrypt(c.file, identity)
	if err != nil {
		return nil, fmt.Errorf("could not decrypt data: %w", err)
	}

	return decrypt, nil
}

func (c *Container) Writer() (io.WriteCloser, error) {
	recipient, err := age.NewScryptRecipient(c.secret)
	if err != nil {
		return nil, fmt.Errorf("could not create crypt recipient: %w", err)
	}

	encrypt, err := age.Encrypt(c.file, recipient)
	if err != nil {
		return nil, fmt.Errorf("could not encrypt data: %w", err)
	}

	return encrypt, nil
}
