package crypto

// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

import (
	"context"

	"github.com/manifoldco/torus-cli/envelope"
	"github.com/manifoldco/torus-cli/identity"
	"github.com/manifoldco/torus-cli/primitive"
)

// SignedPublicKey returns a PublicKeyEnvelope that is signed, and includes
// a tamper-proof ID.
func (e *Engine) SignedPublicKey(ctx context.Context, body *primitive.PublicKey,
	sigID *identity.ID, sigKP *SignatureKeyPair) (*envelope.PublicKey, error) {
	id, sig, err := e.signAndID(ctx, body, sigID, sigKP)
	if err != nil {
		return nil, err
	}

	return &envelope.PublicKey{
		ID:        id,
		Version:   uint8(body.Version()),
		Body:      body,
		Signature: *sig,
	}, nil
}

// SignedPrivateKey returns a PrivateKeyEnvelope that is signed, and includes
// a tamper-proof ID.
func (e *Engine) SignedPrivateKey(ctx context.Context, body *primitive.PrivateKey,
	sigID *identity.ID, sigKP *SignatureKeyPair) (*envelope.PrivateKey, error) {
	id, sig, err := e.signAndID(ctx, body, sigID, sigKP)
	if err != nil {
		return nil, err
	}

	return &envelope.PrivateKey{
		ID:        id,
		Version:   uint8(body.Version()),
		Body:      body,
		Signature: *sig,
	}, nil
}

// SignedClaim returns a ClaimEnvelope that is signed, and includes
// a tamper-proof ID.
func (e *Engine) SignedClaim(ctx context.Context, body *primitive.Claim,
	sigID *identity.ID, sigKP *SignatureKeyPair) (*envelope.Claim, error) {
	id, sig, err := e.signAndID(ctx, body, sigID, sigKP)
	if err != nil {
		return nil, err
	}

	return &envelope.Claim{
		ID:        id,
		Version:   uint8(body.Version()),
		Body:      body,
		Signature: *sig,
	}, nil
}

// SignedKeyringV1 returns a KeyringV1Envelope that is signed, and includes
// a tamper-proof ID.
func (e *Engine) SignedKeyringV1(ctx context.Context, body *primitive.KeyringV1,
	sigID *identity.ID, sigKP *SignatureKeyPair) (*envelope.KeyringV1, error) {
	id, sig, err := e.signAndID(ctx, body, sigID, sigKP)
	if err != nil {
		return nil, err
	}

	return &envelope.KeyringV1{
		ID:        id,
		Version:   uint8(body.Version()),
		Body:      body,
		Signature: *sig,
	}, nil
}

// SignedKeyring returns a KeyringEnvelope that is signed, and includes
// a tamper-proof ID.
func (e *Engine) SignedKeyring(ctx context.Context, body *primitive.Keyring,
	sigID *identity.ID, sigKP *SignatureKeyPair) (*envelope.Keyring, error) {
	id, sig, err := e.signAndID(ctx, body, sigID, sigKP)
	if err != nil {
		return nil, err
	}

	return &envelope.Keyring{
		ID:        id,
		Version:   uint8(body.Version()),
		Body:      body,
		Signature: *sig,
	}, nil
}

// SignedKeyringMemberV1 returns a KeyringMemberV1Envelope that is signed, and includes
// a tamper-proof ID.
func (e *Engine) SignedKeyringMemberV1(ctx context.Context, body *primitive.KeyringMemberV1,
	sigID *identity.ID, sigKP *SignatureKeyPair) (*envelope.KeyringMemberV1, error) {
	id, sig, err := e.signAndID(ctx, body, sigID, sigKP)
	if err != nil {
		return nil, err
	}

	return &envelope.KeyringMemberV1{
		ID:        id,
		Version:   uint8(body.Version()),
		Body:      body,
		Signature: *sig,
	}, nil
}

// SignedKeyringMember returns a KeyringMemberEnvelope that is signed, and includes
// a tamper-proof ID.
func (e *Engine) SignedKeyringMember(ctx context.Context, body *primitive.KeyringMember,
	sigID *identity.ID, sigKP *SignatureKeyPair) (*envelope.KeyringMember, error) {
	id, sig, err := e.signAndID(ctx, body, sigID, sigKP)
	if err != nil {
		return nil, err
	}

	return &envelope.KeyringMember{
		ID:        id,
		Version:   uint8(body.Version()),
		Body:      body,
		Signature: *sig,
	}, nil
}

// SignedCredentialV1 returns a CredentialV1Envelope that is signed, and includes
// a tamper-proof ID.
func (e *Engine) SignedCredentialV1(ctx context.Context, body *primitive.CredentialV1,
	sigID *identity.ID, sigKP *SignatureKeyPair) (*envelope.CredentialV1, error) {
	id, sig, err := e.signAndID(ctx, body, sigID, sigKP)
	if err != nil {
		return nil, err
	}

	return &envelope.CredentialV1{
		ID:        id,
		Version:   uint8(body.Version()),
		Body:      body,
		Signature: *sig,
	}, nil
}

// SignedCredential returns a CredentialEnvelope that is signed, and includes
// a tamper-proof ID.
func (e *Engine) SignedCredential(ctx context.Context, body *primitive.Credential,
	sigID *identity.ID, sigKP *SignatureKeyPair) (*envelope.Credential, error) {
	id, sig, err := e.signAndID(ctx, body, sigID, sigKP)
	if err != nil {
		return nil, err
	}

	return &envelope.Credential{
		ID:        id,
		Version:   uint8(body.Version()),
		Body:      body,
		Signature: *sig,
	}, nil
}

// SignedKeyringMemberClaim returns a KeyringMemberClaimEnvelope that is signed, and includes
// a tamper-proof ID.
func (e *Engine) SignedKeyringMemberClaim(ctx context.Context, body *primitive.KeyringMemberClaim,
	sigID *identity.ID, sigKP *SignatureKeyPair) (*envelope.KeyringMemberClaim, error) {
	id, sig, err := e.signAndID(ctx, body, sigID, sigKP)
	if err != nil {
		return nil, err
	}

	return &envelope.KeyringMemberClaim{
		ID:        id,
		Version:   uint8(body.Version()),
		Body:      body,
		Signature: *sig,
	}, nil
}

// SignedMEKShare returns a MEKShareEnvelope that is signed, and includes
// a tamper-proof ID.
func (e *Engine) SignedMEKShare(ctx context.Context, body *primitive.MEKShare,
	sigID *identity.ID, sigKP *SignatureKeyPair) (*envelope.MEKShare, error) {
	id, sig, err := e.signAndID(ctx, body, sigID, sigKP)
	if err != nil {
		return nil, err
	}

	return &envelope.MEKShare{
		ID:        id,
		Version:   uint8(body.Version()),
		Body:      body,
		Signature: *sig,
	}, nil
}