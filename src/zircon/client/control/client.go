package control

import (
	"errors"
	"zircon/apis"
	"zircon/rpc"
)

type client struct {
	fe    apis.Frontend
	cache rpc.ConnectionCache
}

// Construct a client handler that can provide the apis.Client interface based on a single frontend and a way to connect
// to chunkservers.
// (Note: this frontend will likely be a zircon.frontend.RoundRobin implementation in most cases.)
func ConstructClient(frontend apis.Frontend, conncache rpc.ConnectionCache) (apis.Client, error) {
	return &client{
		fe: frontend,
		cache: conncache,
	}, nil
}

// Allocate a new chunk, all zeroed out. The first write must be done with version=0.
// The chunk is not considered to exist until that first write is performed.
// If this chunk isn't written to before the connection to the server closes, the empty chunk will be deleted.
func (c *client) New() (apis.ChunkNum, error) {
	return 0, errors.New("unimplemented")
}

// Read part or all of the contents of a chunk. offset + length cannot exceed MaxChunkSize.
// Returns the data read and the version of the data read. The version can be used with Write.
// If the chunk does not exist, returns an error.
func (c *client) Read(ref apis.ChunkNum, offset uint32, length uint32) ([]byte, apis.Version, error) {
	return nil, 0, errors.New("unimplemented")
}

// Write part or all of the contents of a chunk. offset + len(data) cannot exceed MaxChunkSize.
// Takes a version; if the version is not AnyVersion and doesn't match the latest version of the chunk, the write is
// rejected.
// Returns the new version, if the request succeeds, or the most recent version number, if the request fails due to
// staleness.
// If the chunk does not exist, returns an error. If this fails for any reason, there must be no visible change to
// the underlying data. If this fails for a reason besides staleness, the version must be zero.
func (c *client) Write(ref apis.ChunkNum, offset uint32, version apis.Version, data []byte) (apis.Version, error) {
	return 0, errors.New("unimplemented")
}

// Destroy a chunk, given a specific version number. Version checking works the same as for Write.
// If the chunk does not exist, returns an error.
func (c *client) Delete(ref apis.ChunkNum, version apis.Version) error {
	return errors.New("unimplemented")
}

// Close all connections used by this client.
func (c *client) Close() error {
	return errors.New("unimplemented")
}
