/*
 * Rosetta
 *
 * A standard for blockchain interaction
 *
 * API version: 1.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api

type MempoolResponse struct {
	TransactionIdentifiers []TransactionIdentifier `json:"transaction_identifiers"`
}