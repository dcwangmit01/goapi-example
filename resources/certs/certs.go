package certs

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

var _insecure_key_pem = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x97\xb5\x0e\xec\x8a\x01\x05\x7b\x7f\xc5\xed\xad\xc8\x4c\x45\x0a\x33\x33\xbb\x33\xad\x79\x0d\x6b\xfe\xfa\xe8\xdd\xa4\xcc\x69\x4f\x37\x9a\x66\xfe\xf5\xcf\x38\x51\x56\xad\x3f\x9e\xcf\xfe\x71\x3c\x35\x62\x03\xf1\x8f\x2e\xa6\x7f\x1f\xc0\x54\x55\x4d\x67\x55\x8e\x65\x75\x9e\x6d\x44\xf6\x9b\x5a\xaf\x1a\x7c\xcf\xaf\x7e\x0a\x57\x1d\x38\x35\xee\x22\x39\xcc\x0b\x68\x4a\x2d\x47\x33\x79\xac\x14\x33\xe5\x81\x8b\x15\xd4\x89\xe2\xe8\x94\x80\xaf\x0a\xa9\xe7\x0c\xfa\x07\xfd\x68\x94\xd3\x20\x2a\xcf\x5e\xe4\x18\x15\x57\xe6\xc9\x82\xee\xff\x5a\x93\x42\x53\xff\x8e\xf5\xdf\x07\x6a\xb3\xdb\xe5\xa3\x50\xbf\xa2\xf3\x31\xc3\x33\x33\x72\x13\x06\xe8\xb1\x37\x15\x4d\x26\x6f\x35\x6b\xf0\xc6\xac\x76\x37\xf1\xae\x9e\x42\xd9\x26\xf3\x2a\x9a\x11\x62\x29\x7e\x52\x51\xbc\x48\x83\xc8\x20\x0d\x2a\xef\x32\xc8\x4a\x7b\xbb\xaa\xd3\x37\x77\x43\x5b\x00\x7d\x18\xf6\xb6\xa3\xaa\x79\x6e\x85\xc1\xb5\xfa\x5a\x65\xd1\x19\x6c\xee\xdc\x4e\xdb\x38\xeb\xfa\x9f\xb7\x84\xdd\x0f\x09\xea\xed\xf7\x79\xa8\xd6\x07\x91\x87\xe9\xb5\xbd\x55\xad\xa5\x92\x8b\x44\x06\xda\x52\xec\xd2\x5b\x3f\x4c\x3e\xe3\x27\x42\xe7\x5a\xfa\xa8\x6e\xae\x67\x40\x08\xd6\xa6\x7d\xd1\x3c\x47\x08\xc5\x3b\x46\x49\xad\xca\x64\x65\xf5\x46\xfb\xea\x96\x81\xda\x58\x72\x11\x6d\x72\xa1\x01\x1d\x52\xe3\x80\x4e\x1e\x49\xaf\x5c\x84\x96\x6f\x23\xad\xde\xa0\x79\xc5\x34\x50\xb5\x77\x3f\x55\x7c\xda\x57\xd1\x5e\xc9\x9d\x36\x07\xfb\xfd\xe5\x4f\x64\x80\xdd\xba\xc4\xda\xd4\x58\x75\xb9\x6f\xc0\x2c\x1b\x11\xde\xbb\xf7\xb9\xc3\x56\xe6\xef\x72\x3c\x2b\x4c\x01\x5b\xa1\x14\xa1\x29\x72\xbd\x09\x9c\xc6\x7d\xe0\x9b\x7a\x21\x0c\x4b\xf5\x7c\x58\x7d\xaa\xe5\xc9\x22\xe8\xd2\x75\x78\x77\x15\x4b\x60\x6d\x13\xff\xad\x82\xc7\xf1\xc6\xeb\x6d\x73\xda\x20\xb2\x7d\x36\x57\x2d\xa4\x58\xb6\x7d\x5f\x7d\xfd\x04\xf8\x0b\x3a\x42\xcb\x6e\xf4\x06\x7f\x1c\xa6\x17\xdb\x20\xcd\xf0\xe6\x4b\x4d\xee\x51\x79\x40\x55\xc8\x93\x44\xd3\x42\xbc\x3e\xf0\x38\xd8\x71\x0b\x0e\x21\x4e\xe3\x0c\x88\xa5\x5f\x99\x09\x43\xc4\x0d\x55\xfd\xa4\x30\x56\x5b\xa7\xb5\xb0\x28\x0e\x85\x4f\x16\x3f\x0f\x62\x30\x6a\xdb\x79\x14\xe0\x29\x3a\x08\xed\x55\x6c\x24\x96\x8c\x5c\x73\x32\x6a\x8b\x66\xe0\x38\x4c\x6f\xf9\x1b\x72\xb4\x6f\x9a\x7d\x5c\x17\xa6\xc1\x4f\xa0\x87\x0e\x3a\xfe\x64\x25\x77\x44\xc8\x5a\x6f\x29\xda\xfd\x04\x7a\x00\xc3\x27\x48\x48\xd3\x84\x02\xfc\xe8\xe5\xdc\x7e\x97\x59\x85\xee\x69\x9e\xd8\xf1\x13\xd4\x41\x4f\x69\x6f\xb6\x38\x60\xea\x6c\xf7\xa6\x6d\x51\xb7\x67\x0e\xa5\x19\x6d\xd5\x20\xac\xc8\xb3\x97\xc8\x02\xac\xfb\x5f\x85\x07\x6e\x7e\x63\x57\x62\xd8\xd1\x3a\x74\xc4\xad\x7f\xc5\x43\x54\x23\x89\x36\x89\xba\x3e\x62\x44\x3a\xcc\x70\x8b\x6d\xa4\x79\xf0\x83\x61\x28\x1f\x44\x03\xf2\x08\x23\x54\x02\x37\xed\x61\xfe\x29\x8f\xac\xd1\xf8\xab\x00\xb5\xfe\xc5\xb8\xeb\x46\x10\xfc\x6a\xb7\x79\x6f\x92\x67\x51\xec\xae\xb7\x68\xf9\xce\x42\x5f\xaa\x7d\x7f\x05\x94\x54\x4c\xba\xc4\x91\xc2\xcc\x93\x44\x03\x4d\x25\xf2\x3d\x69\x86\x7a\x3a\x75\xd9\xd8\x56\xfd\xb4\xf6\x31\x85\x8d\x47\xf9\x1d\x59\xd6\x17\x0d\xe4\x0b\x06\x38\x5c\xc5\x12\xaa\x6a\x86\x41\x9d\x69\xbb\xed\xb5\xa0\x18\x64\x84\x28\x19\xc9\x02\x2b\xa4\xab\x5c\xa3\xa6\xdf\xdc\x97\x51\x6d\x0f\x66\x73\xbc\x23\xe4\x20\xb1\x66\xab\x36\x5c\xa7\x82\xe2\x9a\xc4\x3c\x4e\x92\x68\x7c\x14\x49\x95\x67\xef\x95\xcc\x46\x52\xa5\xdb\xed\xa3\x79\x86\x80\x8e\xb1\x06\x70\x19\x46\xb5\xbc\xb2\xc1\xa4\x9a\xb6\xf9\x5e\xaa\xd6\x22\x43\xf1\x8d\x40\xdc\xf1\x74\xb2\x0a\x94\x2f\x45\xd1\x70\x4d\x78\x58\xa5\xfd\x8c\xfa\x25\xce\x4c\x45\xb2\xda\x20\x21\x54\x01\x3c\xe5\xf3\x88\xae\x3c\xf1\x64\x7a\x3d\x4d\xaa\x2b\x58\x4c\x63\x11\x54\x9c\x94\xe4\x7c\x72\xb8\x2b\xc1\xa8\x14\x48\x6e\x25\xa8\x24\xa1\x88\x23\xc9\x46\x38\xf8\x9c\x78\x85\x8c\xaf\x14\x5f\x81\x0d\xa0\x0f\x1b\xac\x49\x0c\x45\xcd\xc9\x6b\x8f\xcc\x34\x2c\x83\x56\x58\x2e\x3d\x74\xa8\xe7\xfc\xfc\x10\x3c\x72\x06\x97\x43\xc7\x0b\x58\x32\x97\x64\xbb\x9f\xd1\xff\xac\x1c\x59\x05\xe2\x39\xcf\xfa\x17\xd0\xdd\x9c\x87\x3e\x1b\x1f\x65\x53\x60\x98\xaf\x01\x1a\xc6\xf1\xa6\x7a\xac\xd3\x7a\xa1\x36\x94\x6c\x67\xd8\x6b\x31\x8f\x67\xb1\x7b\x2f\x72\x8c\x19\x48\x70\x72\xbc\x3b\xf8\xec\xed\xf4\x41\xd7\x1a\xa8\x76\x57\x9e\x0d\xd9\xc2\x38\xad\x6d\x54\x37\x82\x33\x1c\xe9\xfb\xf4\x21\x73\xff\xa4\x8e\xb8\x0c\xd9\xc1\xd7\x5c\x82\xa2\x2e\x3f\x4b\xc2\x5f\x98\x3f\x54\xaf\xa6\x9f\x37\x8e\x18\xe8\x08\x5a\x01\x88\xf9\x4e\x7a\xb9\xdc\xe8\x58\xe5\xcc\x0f\x48\x39\x87\x49\xca\x42\xef\x90\x52\x22\xc2\x99\x83\x19\x18\x04\xd4\x63\x53\xae\x8e\x88\x06\x0d\xfd\x89\xa4\xa1\xd7\x6c\x05\x14\xc2\x5c\xa5\xfb\xc1\x07\x1c\x13\x5a\x38\x81\xdd\x31\xe3\xad\xc2\x1c\x62\x52\x4e\x7a\x45\xe8\xa1\xf8\x1e\x42\x8b\xed\x9c\x39\xbf\x4c\x5d\xa2\xa4\x7a\x38\xaf\x3d\x6d\x9d\xee\x23\xe4\x9b\x46\xe4\x58\x73\xc6\xef\x8f\x7e\x02\xed\xfe\x66\x55\x27\x45\x99\x2d\x64\x32\x1c\x83\x41\xaf\x43\x6a\x87\xd3\xcc\x57\x60\x7c\x66\xd4\x2e\x5a\x89\xb7\x51\xe3\x7a\xee\x52\x9c\xa5\xac\x3c\x39\x58\x06\x74\xd8\x7b\x82\x32\x54\xd9\x10\x01\x2f\xa8\xe1\xc5\xb9\x87\xee\x35\x14\x97\xe9\xa2\xf3\x77\x55\x14\x75\xb7\x06\x7c\x91\x28\x41\x45\x72\xd5\x95\x51\x43\x75\x13\x9b\x60\xa0\x68\x9c\x5f\x46\x91\x04\x17\x4f\xcf\x2a\x2d\x4c\x70\x89\x81\x02\xda\x21\x81\xfc\x55\x58\x5c\x2f\xec\xed\x5c\x50\xe9\xd4\x65\x0f\x0a\xbb\xfd\xc6\xe1\xb9\x41\x43\xa1\x3e\xa9\x6a\x14\x4e\x1c\x7b\x84\x11\xca\xa8\xe5\x98\x24\x46\xb7\x3d\xde\xb1\x5c\xcf\x7c\x00\x43\x81\xe2\xbd\x45\x62\x5e\x1f\x8c\x8b\xc2\xea\xbb\xbd\x2d\x6a\x8d\xbc\x4a\x02\x57\x08\x22\x97\x37\x58\xa1\x09\x16\xc8\x6d\x4e\x24\xa6\x10\x94\xd5\xc8\x24\x61\x18\xa1\x25\x3e\xc1\xb0\x0f\x65\xe0\x5b\xc6\xd6\x05\x49\xc7\x53\xd9\x2a\x7a\xf2\x30\x92\x0a\x70\xaa\xde\xe2\x70\x08\x89\xa6\xaf\x47\x32\x71\xfb\xe7\x9d\xbc\xe2\xb0\x23\x67\x68\x1b\xa5\xf4\x4c\xfe\x86\x43\x1e\xb5\x3f\x70\x0e\x17\x80\x88\x40\xe1\x27\x6b\x73\x87\x9c\xd1\x9c\x31\xff\x47\xb8\x9a\xe8\x28\x1c\x2b\x1c\x2c\x82\x3b\x0a\x85\x3a\xeb\x09\x36\x37\x6c\x1e\xc6\x41\x34\x4b\xe3\xc3\x89\x6f\x65\xee\x47\x87\xbd\x80\xbe\x5a\xf3\x52\x52\x62\x4b\x1b\x60\x1f\xb9\xee\x52\xc8\xe8\x9b\xda\x59\x08\x47\x8b\x52\x8f\xb2\xc0\x6a\xb9\xd1\xdf\xf3\x7d\x9a\x74\xe3\x6d\xac\x2c\x4a\x2b\x7b\xdb\xdf\x44\xb1\xf2\x91\x1e\x30\xe0\x65\x31\x82\x29\xaf\xbd\x60\xef\x55\x84\xbb\x64\x7a\xc8\x7a\xb7\x7f\x0e\xe7\xfa\x55\xd5\x28\x26\x49\x62\xd2\x1b\x9b\xac\xca\xb0\x74\xc8\x6e\x19\x65\x1d\x57\x66\xc5\x7e\xdd\x1a\x7d\x31\xba\x3a\x30\xc1\xe1\x37\x2e\x62\x7f\xc9\x3e\xcc\x6c\x09\x6f\x77\x15\x70\x53\x1e\x06\x63\x27\xb7\xf1\xe2\xa6\x84\x34\xe2\x6a\x7f\x92\x14\xe6\x46\xb1\x9d\xe7\x4f\x38\x2d\xca\xde\x1b\x32\x91\xf2\xb9\x60\xa9\x40\x67\x2b\xc6\x79\x33\x53\x60\x3f\x0b\xd8\xb8\x7b\x1b\xde\x0d\xd8\x13\x41\xb7\x14\xd5\x49\x9b\xed\xdb\x9a\x38\x4f\xe6\xf4\xfd\xa1\x82\x64\x06\x29\x43\xb6\xbe\x19\x01\xc7\x1b\x95\x67\x64\x64\x35\x00\x4d\xc6\x2d\x38\x6f\xcd\x96\x46\x47\x7c\x79\x8e\x6c\x78\x39\x89\xd4\xce\x51\xb8\xdc\x5d\x65\x26\xf6\xc0\xb2\x87\x97\x52\x23\xfd\xfe\x01\xcc\xea\xeb\xbc\x5f\x51\x58\x19\x2c\xa2\xe5\x9d\xc0\x00\xda\xc2\x28\x9b\xc9\xa9\x60\x94\x5c\x4a\x50\x69\x35\xca\x48\x9d\x62\xfb\xbb\x76\x66\x5b\xfa\x2b\x6f\x4c\x0c\x50\xb5\xc9\x1c\x07\x46\x11\xca\x64\x55\x1f\x0d\x79\x5d\x97\x3d\x76\x64\xf0\x9c\x50\x01\x8c\xa5\xc4\x71\xcb\x35\xbf\x94\x36\xbd\xe3\x75\xc8\xab\x3c\xca\x54\x7a\xd3\xdb\x7d\xdc\x32\x13\xe2\xa3\x91\x21\x0e\xb3\x20\x87\x5b\x74\xab\x5d\xee\x69\xd0\x61\x83\xce\xa3\xe1\x6e\xd2\x7d\xa1\x00\x1d\x50\xbf\x74\x94\x4f\x5e\xfb\x9d\x72\x8a\xec\x1a\x59\x84\x89\x97\xd1\x33\x4d\x07\x2a\xd9\xf5\x98\x2a\xe6\xe5\xcf\xac\x85\x76\x9a\xc0\xfa\x49\xce\x43\x29\x28\x22\x68\x49\x62\xab\x43\xcc\x5d\x81\x34\x6b\x42\x24\xb4\x3a\xf5\xe8\xba\x98\xed\xe5\x8a\xf6\x09\x1b\x22\x3d\x3f\xd8\xbc\x8b\x0a\x1a\x1e\xb5\x94\xab\xb6\x0e\xfe\x43\x2e\x08\xab\x20\x69\xa9\x2b\x3d\x41\xc0\xa7\xc7\x91\x54\x41\x48\x80\x6a\x94\xb8\xd8\x52\x23\x6f\xd2\x72\x4e\xc9\x97\x84\x25\xbf\x07\x57\xe8\x32\x49\xd4\xd5\x9a\x6e\x67\x79\xbf\x5e\xfb\xc9\x02\xb0\x51\x22\xe1\x62\xb2\xa3\x25\xcd\x84\x34\x48\x3b\xd4\xa4\x45\xa7\x00\xb2\x5e\x9d\xea\xd2\x79\xd6\xe5\x76\xb0\x2f\x4d\x38\xa8\x72\xfb\x7d\x66\xc6\x84\xab\xf7\xc8\x0c\x27\xff\x0e\x16\x62\xb5\x5a\x98\x6e\x0c\x81\x3f\x1c\xac\xc0\x0a\x65\x31\xb9\xe8\xd7\xe7\xd8\xb8\x40\x40\x8e\x33\x52\xd9\x8d\x5c\x10\x4a\x76\x37\x06\x2c\x4c\x31\x6a\x3b\x9b\xfd\xe5\x99\x61\xe3\x2b\xd1\x3b\x77\x48\x8e\x6e\x18\x1b\x28\xa7\xfe\x21\x9b\x50\x7f\xdc\x78\x50\x35\x89\x76\x24\x26\xd3\x01\x79\xa6\x69\x53\x51\x86\x6e\x65\x73\xc3\xf3\xa5\xd2\xdb\x6d\xea\xe8\x47\xc4\x79\x7f\x0b\x06\x67\xe6\x6f\x5d\xe1\xe4\xf7\xde\xcc\x85\x54\x0f\xd1\x8a\x9a\xf9\x71\x1f\x2d\xf4\x98\xaf\xbb\xc5\xbf\x0d\x98\xdf\x88\xf0\x93\x6f\xde\x74\x0f\xed\x7b\xdb\x99\x30\x02\x3b\x36\x79\xc7\xde\xd0\xde\x9f\xd7\x0f\x6b\x95\x04\x23\x44\xbd\x4e\x48\x83\x22\xbf\x7b\x6c\xe8\x02\xf1\xb5\xe7\x25\x1f\x17\xa6\xe5\x7a\x20\xc0\xfc\x64\x0b\x8f\xf6\xdd\x0c\xb4\x09\x86\x81\x56\x02\xa6\x36\xd0\x4e\x82\x10\x49\x81\x03\xf3\x8a\xed\xd2\x99\x31\x1e\xcb\xf6\x49\xfb\x95\x8e\x12\x73\x88\xd2\xcf\xc2\x38\xc3\x3a\xc6\xf8\x21\xd0\x87\x66\x8b\xc3\xe9\x8a\x3f\x1a\x28\xfc\x22\x63\xb1\xf2\x15\x49\xcb\x0c\xe3\x0b\xc1\x20\xed\x88\x9d\x55\x8e\x55\x76\xe7\xf7\x34\xe4\xc6\x93\xa7\xe0\xf3\x8b\xaf\xec\x65\x8a\xa3\xe8\x08\xe1\x20\x20\xc6\xc5\x08\xbe\x96\x44\xab\xaa\xc3\x5a\xc8\x41\xe8\xaa\xfd\x23\xb4\xa2\x82\xcc\x79\x60\xf3\x06\x3c\x62\xd5\x1a\x16\xc6\x9f\xf3\xe6\x4e\x59\xda\x3e\xa4\x6d\xb1\x8f\x96\x55\x4d\x0d\x2e\xa8\x09\x90\x8e\xc7\x20\x61\xee\x77\x96\xe2\x78\x4e\x26\x65\x56\x67\x14\xa5\x85\xc5\x33\x15\x6c\x7c\x37\x92\x46\xbe\x76\x7d\x20\x23\xac\x98\xf5\x2c\x3d\x62\xa2\x77\xbc\x99\xe4\xb7\xb9\x63\x61\x6a\x08\x25\x70\x14\xeb\x44\x70\x4d\x0b\x59\xc1\xa5\xda\x14\xa9\xbc\x81\xb8\x7a\x8e\x0b\xee\x07\x35\x31\x59\xaf\x39\x53\x53\x82\x08\xe9\xf3\x31\xb5\xa8\x5e\x6e\xe1\x29\x6f\x32\x42\xde\x49\xfa\x6c\x61\xcf\x09\xb0\xa6\xc8\x61\xeb\x8c\x39\xd1\xf7\x98\x3f\xb4\x7e\x17\x61\xb9\x04\xb2\xdb\x57\x58\x77\x1e\xa2\x54\x99\x73\xe4\x8c\x89\xd3\x3e\xce\xab\x9e\x8c\xce\x74\x2f\x49\xf2\x4d\xc6\x7e\xe7\x33\xaf\xb5\x04\x98\xa5\xb4\x58\x11\xef\x50\x78\x56\x75\xd8\x5a\x29\x18\xda\x37\x10\xc4\xd5\x9b\x84\x4a\xd8\x00\xca\x17\x28\x3f\xbf\x67\x71\x35\x82\x4a\x9e\x24\xba\x9e\xe1\x66\xb3\x6c\xf0\x9b\xca\x7f\x03\x7f\x93\x43\xb4\x84\xff\x9f\x22\xff\x09\x00\x00\xff\xff\xe8\x4a\xb6\x6f\xab\x0c\x00\x00")

func insecure_key_pem() ([]byte, error) {
	return bindata_read(
		_insecure_key_pem,
		"insecure-key.pem",
	)
}

var _insecure_csr = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\xd5\xb7\xd2\xb3\x30\x16\xc6\xf1\x9e\xab\xf8\x7a\x66\x87\x64\x6c\x28\x05\x12\xd1\x60\x84\xc9\x1d\x60\x93\xa3\x03\x32\x5c\xfd\xce\xbe\xfd\x9e\xf6\x77\x66\x9e\xf2\xff\x9f\xff\x9d\x82\x74\xd3\xfd\xa7\x22\x3f\x30\x35\x53\x05\x01\xfa\xe7\x23\x1c\xa2\x7b\xf0\xa7\x94\x63\x9a\xe8\x0c\x55\x15\x7c\x80\x0a\x30\x20\x75\x6e\xfe\xd4\x03\x58\x4a\xed\x46\x0a\x48\x03\x30\x44\x81\xe3\x3b\x04\xe1\x14\x46\x18\x9b\x88\x2c\x30\x8d\x7f\x4b\x36\xca\x7b\x31\x0e\x0d\xe5\xf8\x29\xd1\xc0\x1f\x1a\x88\x70\x41\x1a\x9f\x6a\x7f\xb2\x9a\x62\x74\x97\x92\x77\x37\xc7\x67\x89\x4e\xfe\xdc\x46\x3f\xdf\x2a\x26\x77\x48\x85\x68\xcf\xee\x0a\xcc\x12\x8b\xa5\xf2\x38\x5b\x52\x5e\x63\xb3\x00\x7d\x1c\xf5\xad\x03\x2e\x44\xea\xcf\xe9\xfd\x51\xde\x4d\x4d\xd9\xf3\x24\x6b\x1e\x7a\x58\x07\xba\xdc\xa5\x31\xa9\x31\x2f\x7f\x8a\x24\xfa\xe6\xb1\xdb\x3c\xf4\x61\xa3\x8a\x56\xf1\x8a\xf1\x27\x3a\x7e\x4d\xb4\xfa\x6f\x09\x22\x22\xbf\x0b\xde\x6d\x0a\xbd\xd9\x4a\x01\x9f\x9d\x00\x90\x1b\x04\xa4\xae\xcd\xd6\x01\xac\xae\xde\x57\xfd\x6e\x16\x02\xc4\x48\xa1\x00\x0e\x01\x38\x99\x2a\x24\x7f\x0f\x36\x98\x4d\x15\x60\xf5\x51\x0b\x4e\x73\x7b\xd2\xe5\x2a\x79\x8a\xe8\xc8\x97\x06\x46\x2f\x6c\x7a\x59\xf7\x1d\xcf\x6a\x86\x40\x54\x70\xfb\xa5\xf6\x05\x8a\x36\xf1\x3d\x94\x7f\x66\xed\x36\x88\xc6\x0b\x93\x49\xd3\x9b\x86\x61\x6d\x42\xb8\x46\xd1\xa6\x0c\x28\x75\x16\x7b\x47\xc7\x96\x35\x57\x2d\x5d\x7f\x0c\xad\x4d\x1a\xb0\x22\xfc\x52\x62\xc6\xfc\x51\x77\xb9\xff\xcc\x47\xb0\xc7\x0e\x29\xfb\xe2\x85\x9a\x11\xb6\xe0\xc7\x73\xca\xc3\x50\xbd\x2d\x55\x95\x41\x13\x0e\xd6\x4d\xa3\xc2\xaa\x71\x00\xe7\x6d\x98\xa4\x49\x3e\xb4\x9f\x3b\x90\xf3\xa6\xf0\x94\x6c\xee\xec\x7b\xec\xd7\xfb\x9d\xd7\xaf\x1f\xbe\x6d\xd1\xcd\x07\xbc\x7b\xaa\x61\xde\x9d\xa1\x36\x8b\xad\x96\x8d\xf6\xed\xa7\x0f\x77\xfb\x88\xdc\xf3\xf7\xc7\x5c\xbf\xda\xf5\x92\x5c\x79\x67\xe4\x51\x4d\xf1\x03\x9b\x7d\xca\x5c\x23\x77\xab\x43\xeb\x89\x58\xbd\x5d\xf4\xb3\x6e\xec\xa2\x80\x74\x4f\xb8\x30\x78\xcc\xf9\x7e\x40\x92\x8b\x03\xad\x28\xe6\x81\xeb\xd3\xe7\xa2\x87\x17\xd5\x1e\x83\xef\x1b\x52\x6b\x8f\x2f\xeb\x44\x5e\x52\x93\x79\xbf\xca\x0c\x17\x0e\x26\xa4\x40\x9f\x46\x70\x6f\xd0\xf1\x1b\xa7\x9e\x6e\xbc\xd4\xb5\x39\x3f\xe3\xf3\x6b\x15\xc8\x8e\x2f\xb7\xdf\xcb\x0a\xb7\xb3\xbd\x0e\xe9\x48\xe5\x80\x13\x77\xfe\x5b\xa7\x9f\xa4\x75\xa1\x4e\x0b\x98\x1b\x6c\x36\xcb\xeb\x8a\xdb\x30\x87\xa3\x24\xeb\x12\xe8\x56\x7e\x93\xbb\xae\x49\x96\xf7\xed\x3e\x08\x96\x04\xd6\xec\x94\xa7\xc2\xf4\xfb\x51\x40\xb2\xd4\xd9\xdf\x57\xed\x71\x75\x39\xcf\x94\xf5\x04\xde\xb4\x17\xd9\x86\x91\xaf\xf7\x79\x08\xde\x40\x35\x3c\x67\x5d\x64\xaf\xf3\x4e\xd2\x0d\xa9\xdb\xbe\xf9\x12\xc3\x3b\x77\xcd\x6d\xa6\x56\xa5\xaa\x6f\xa6\x8a\x9c\xc6\xbd\xf3\x34\x39\x88\x9b\xbf\xae\xf7\x18\x8b\xb9\x71\xf6\x83\xee\xe8\x84\xcb\xa3\x2d\x0b\xce\x0f\x22\x05\xd7\x2b\xb3\x95\x60\x5c\xf3\xe5\x2d\xbc\x8d\xe2\x2e\xc1\x96\xbe\x8c\x14\xfe\x9c\x4e\x8c\x59\x59\xdf\x2b\x93\x3b\xdd\x33\x19\x87\x69\x5d\x5e\x4b\xac\x7d\x55\x8c\x8c\x72\x45\x57\x43\x92\xc6\x47\x33\xc8\x0a\x0c\x0e\x51\x76\x16\xcc\xe4\x31\x49\x1f\xb3\x14\x48\xfc\x0b\x51\x51\xcc\x59\x15\x63\x7e\xae\xd3\x4b\xea\x7b\x77\x13\xa5\x65\x6f\xf5\xe7\x58\x77\x8c\x9e\xb7\xa9\x1a\xc9\xae\x78\xf3\xde\x93\x3b\xf6\x8c\x58\xd3\x9b\xfe\x1e\x3f\xb1\xcd\xf5\xcc\xbb\xff\x6a\x3c\x0c\x29\x80\x4d\x08\x30\x50\x66\x00\x08\xc4\xa9\x65\xcf\x99\xd9\x6c\xa5\x0b\x30\x72\x15\x0c\x60\x5d\x9b\x0a\xd0\xbf\x2f\x30\x73\xb3\xde\xbc\xaa\x4b\xb5\x38\xf4\xe9\x65\xaf\xf4\xd4\x29\xd9\x49\xa2\xbe\x71\x3d\xd1\x95\x91\xef\xdc\x94\xf2\x8d\x55\x32\x86\x0d\x7d\xf6\xe8\x8d\x44\x10\xfc\x47\xbf\xd0\x4b\xf5\x8b\x88\xf6\xa9\x34\xc8\xc0\x6c\x2e\x57\xa3\x68\xe3\x0b\xf9\xb0\x8e\xb4\x47\x6c\x5f\x52\x83\xed\xda\xc9\x40\x64\x40\xe4\x3d\x65\x91\xc1\x38\x03\x09\xa3\x08\x1f\x34\x57\xf4\xf8\x1c\xc7\x2a\x0b\xfa\x0b\xed\xbb\xe4\xb1\x1b\xc5\x4f\xad\x9c\x15\x17\x91\xed\x6b\x6d\xbd\xf8\x4e\x5d\x50\x8f\xeb\x66\xf2\xe3\x92\x2d\xdc\x60\x76\x10\x26\xe1\x24\x88\x6c\x07\x0d\x06\xb0\x8d\xc6\x1d\xaa\x0b\xf4\xd5\x4c\xc6\xb7\xbd\x9d\x2d\xe4\x09\x82\xbc\x0b\xa6\xbb\x7d\x1d\xfa\xf6\x4c\x3e\xb4\x6c\x51\x48\x6b\x14\x85\x78\x07\x44\xbd\x52\xdc\x3e\xe7\xd7\xa1\xe8\x33\x31\xe1\xca\x75\xcd\xa1\x5d\x98\x9b\xfe\x76\x79\xdf\x93\x68\x6e\x30\x58\x5e\x49\x98\x31\x5a\x2c\x90\x6c\x27\x08\xb5\x2b\x67\x03\xea\x53\x96\x9e\xe4\x9b\xdb\x49\x7a\xda\x61\xf3\x55\xf0\xbd\xd5\xa1\xee\xa2\x63\xcf\x91\x00\x73\x36\x48\xde\xcd\x5e\xb0\xca\x25\x62\x56\xde\x5e\xcf\xa7\xc7\x5e\xf8\x0f\x2c\x39\x4f\xce\xbe\x54\xe3\x41\x0d\xc3\xd6\x94\x57\x2f\x74\x64\x42\x8b\xbc\xde\x85\xd1\xc9\x6e\x1b\x9a\xe1\x4a\xfe\x3b\xc2\x8d\x77\xd7\x4b\x15\x4d\x80\xd3\xde\xa0\x70\x53\xad\x1b\x16\x2f\xb9\x94\x4b\x76\x1b\xc8\x8d\x7e\x2a\x1e\x55\x3f\x50\x62\xdd\x32\x5b\xef\x25\x2f\x98\x9e\x50\xb0\xd4\xbb\xba\xf8\xee\xe3\x71\xd4\x55\xf1\xad\x80\x1f\xbf\xb9\x85\x78\xb7\x6e\x9f\xe8\x72\xbd\x72\xcd\x2c\xf7\x7c\xcc\xeb\x6d\x9f\xcb\x31\xb3\x52\xe0\x6d\xb6\xf2\x69\xa1\xe5\xb1\xd9\x1f\x07\x11\xb7\x2b\x3b\x7f\xdd\x83\x68\x24\x0b\xd2\x96\xe9\x58\xb3\x9f\xa7\xc9\x5b\xcf\x31\x7b\x37\x38\x35\x48\x2e\xb7\x5e\x01\x8e\xb1\x93\xab\x12\x1b\xa7\x82\xb2\xab\xc4\x2a\xb7\xb3\x2a\x01\x7c\x0f\xad\x2c\x2e\xa2\x10\x5b\xb6\x72\x0e\x2a\x06\xbf\xb6\x6d\xe7\xf0\x47\x81\x2a\x47\x7f\x5b\x1c\x22\xc7\xe7\xae\x15\x1a\xe7\x07\x5a\x83\xef\xaa\x73\xad\x3b\x53\x92\xca\x36\x43\x51\xaf\x5e\x28\xf2\xb6\xd2\x57\xd6\x28\x86\x43\x29\x3e\x5f\x9f\x9f\x34\xba\xcf\xf9\xb8\x33\x86\x58\x9a\xf8\x79\x65\xbc\xd1\x94\xcf\x27\xb9\x16\x8f\x07\x6c\x3e\x63\x3d\xe2\x42\x7c\x51\x22\x33\xac\xf1\xab\x93\x38\x73\xdc\x1d\x72\x11\xa8\xbf\x60\x21\x17\xfe\xff\x98\xfd\x37\x00\x00\xff\xff\x0f\x4b\x37\xce\xf1\x06\x00\x00")

func insecure_csr() ([]byte, error) {
	return bindata_read(
		_insecure_csr,
		"insecure.csr",
	)
}

var _insecure_pem = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x96\xc7\xb2\x83\x3e\xb2\xc6\xf7\x3c\xc5\xdd\xbb\x6e\x99\x68\xc3\x52\x42\x22\xd9\x60\x72\xda\x11\x4c\x32\xd1\x18\xcb\xf8\xe9\xa7\xce\x39\xff\x59\xcc\x8c\x96\xfd\xb5\xaa\xa5\xaf\xfa\xd7\xd5\xff\xff\x73\x20\x56\x75\xeb\xff\x64\xec\xfa\xba\xa2\xcb\xc0\xc7\xbf\x51\xca\xd4\x75\xd5\xe8\x64\x19\x82\x53\x0d\x88\x0e\x41\xad\xeb\x69\x44\x07\xd6\x9b\xe7\xd3\x0a\x13\xe4\x24\xc6\x65\x4a\xf5\xe6\x5d\x58\xc0\xc1\x16\x74\x00\xa9\x33\xfd\x23\x7f\x81\x01\x6b\x2b\xa4\x20\x48\x7c\xd0\x87\xbe\xe9\x9a\x04\x3b\x09\x0a\x1d\x47\xc7\x64\x46\x49\xf4\x99\xd3\x41\xda\xf3\xa1\x6f\x4c\x37\x21\x0a\xf8\xd5\x34\x4c\x18\x3f\x89\xf8\xda\x1d\x8d\x26\x1f\xac\x99\x2a\x58\xeb\x6d\xba\x34\x51\xc9\x6f\xc2\x05\x7f\x5c\x23\x1f\xad\x3e\xe1\xc2\x3d\xf5\x20\x4a\x63\x83\xce\xa2\x74\x4e\x58\x85\x4e\x7d\xfc\x32\xe5\x55\x05\x4c\x80\xe5\x8f\xf9\x70\x07\x69\xa7\x74\x05\xee\x59\x9c\x36\xa5\x1a\xd4\xbe\x2a\x75\x49\x44\x6a\x87\x95\x5e\x79\x1c\x6e\x59\x64\x35\xa5\xda\xbf\xf3\x16\xda\xf9\xf0\x11\x4c\xb7\x26\x4a\xfd\x5b\x08\x61\x22\xad\x39\x6b\x35\x54\xae\x36\xef\x82\x73\x4e\xa6\x0f\xc8\x0d\x01\xa2\x35\x85\x65\x76\x80\x98\x3e\xe0\x4d\xbf\xde\xcd\xaf\xc9\x47\xff\xc4\x2c\x04\x04\x13\x81\xfd\x86\x4c\x3e\xea\x64\x38\x75\xf8\x4a\x99\xe0\xf1\xfb\x24\xd8\x98\x72\x18\x9a\x1f\xfc\x05\xee\x8f\x39\x10\xd4\xbe\xfc\xb0\x9a\x5c\xed\x87\x9c\x33\xb6\x2c\xc2\x1f\xa5\x03\xc1\x9f\x56\xf8\x28\xb4\x9a\xbc\x85\x2a\x55\x0c\xca\x96\xb0\xfd\x37\x61\xc5\x8f\xe6\x83\xfc\x2f\x61\xf2\x15\xdc\x6f\x05\x1b\x76\x65\x6c\xf4\x3a\xb6\xfa\x62\x74\xe7\x74\xe8\xbb\x24\x76\x7b\xd3\xa3\xc9\xe5\xcf\xb3\x2b\x85\x77\x57\xcd\x39\xbd\x0e\x34\x63\x2e\x07\x85\x4e\x3d\x68\xfe\xfc\x2d\x97\x21\xca\x59\xe6\x55\x46\xc2\xaf\x81\x59\x24\x6d\x3a\x96\xb6\x5c\x7b\x7c\x54\x04\xa2\xbf\x42\xa6\x4f\x21\xf6\xf3\x4e\x58\x65\xcd\x54\xe9\x5b\xa2\xe9\x63\x22\xc0\x9b\x48\x96\x41\xab\xff\x77\x0f\x40\xe8\x00\x54\xd7\xba\x0d\x7e\xf4\x7a\x92\xeb\x5a\x87\x14\x30\x58\x54\xec\x58\x38\x0b\xfb\x93\x88\xda\xe3\xcb\x6d\xd8\x8a\x5e\xa0\x96\x86\xdb\x29\x9f\x8c\x87\x03\xc3\x77\x7c\xdd\xa0\x56\xf1\x0d\x34\x7c\x0e\xb7\x80\xc5\xbe\x20\x8f\xd2\xb8\xdf\xbf\x3c\x65\x79\xa0\x36\x6f\x73\x18\x89\x6a\xe0\xd6\xfb\xe3\xa9\xfb\x37\x69\xc0\x6b\xd8\x3e\xb5\x23\x1f\x15\x01\xe8\x43\x79\x55\x72\xb1\x53\xae\xac\xc7\x76\x96\x6d\x24\x5f\xb8\xbb\xdb\xea\xa9\xc9\x4d\x47\x94\x96\x07\x0a\x5d\xe8\x07\xa9\xd6\xd5\x20\xae\x6c\x87\x69\xc2\x75\x80\xd0\xbc\xb5\x87\xa8\xda\x2b\xce\x8c\x31\x1b\xa3\x45\xd4\x72\xba\xbb\x7a\x7b\xe6\xcd\xc6\x35\x49\x36\x2e\xbb\xe8\x8e\x80\x11\x95\x70\x2d\xb0\x16\x7b\x32\xa3\x6e\xd0\xc3\x21\x99\xce\x38\x0b\xf5\xa7\x15\x70\xcb\x59\x13\x37\x3e\x78\xbf\x8a\x77\xb2\xe7\x89\x87\xb2\xd8\x1d\x98\xcf\x14\x43\x7d\x30\xbd\x67\x07\x06\x67\xde\x3c\xaa\x4e\xaa\xeb\x58\x38\xc9\xb1\x7a\x4b\x72\xfa\xcc\xbc\xc0\x27\x0c\x34\xc3\xd7\xd6\xc6\x91\xdb\x9c\x1e\xa9\xbf\xea\x4b\x7a\x3b\x91\xdb\xe2\xa2\x6d\xa9\xe4\xf7\xae\x3c\x8e\x8a\xd8\x78\x63\x60\x55\x90\x5a\x3d\x56\x2b\x68\xde\xfc\xa8\x78\x97\x8b\x73\xbe\xdf\x8c\x67\xd6\xa0\x65\x7b\x56\xc8\x40\xeb\x59\x5d\x7b\xef\x38\x2d\x4b\xd8\xa6\x13\x8a\xc7\x6b\x76\x82\x2d\x73\xd7\x69\x33\x3f\x97\x28\x0a\x9e\x94\xdb\x2f\xf0\x18\x49\x28\x84\x61\x39\x58\x85\xc9\x48\xaa\x32\xd1\x74\x27\x0f\x44\xb8\xc4\xc5\x48\xe4\x79\x6c\xe6\xee\x5e\x69\x18\x91\xc7\xa5\xd1\x2e\x53\x48\xaf\x5c\x20\x76\x74\x5a\x98\x7c\x44\xbd\xe5\x43\x94\x67\xe8\xd2\x86\x37\x02\xf4\x42\xdc\x97\x91\x3e\x98\xc7\x8e\xf0\xce\xe5\x78\x91\xb4\xef\x31\xf9\xe8\x01\xab\xde\x75\xe3\x20\x3c\xae\x63\x10\x47\x9f\xb9\xa9\x6c\xc0\x2c\xeb\xcb\x48\x11\xd5\x4f\xd5\xac\xdc\xec\x9b\x5d\xbd\x58\xe3\xf3\x0e\x15\x2b\x50\x64\xf9\x79\x90\x88\x9c\xcd\xcb\xf0\xbd\x93\xf2\x75\x25\xb7\xeb\x79\x4b\x65\xae\xeb\xc4\x46\x1c\xf8\xb7\x34\xed\x96\x70\xcf\xa2\xfb\x32\x50\x5b\x9f\x44\xbc\x01\x9c\x72\x9f\x9c\xb5\xfa\xee\x29\xab\xc6\x34\xb6\x6c\x7b\xa4\xf7\x1e\x49\x73\x0e\x1b\xb6\xfb\xd8\xdf\x6c\x75\xc3\x3c\xe8\x8f\x62\x4b\x6f\xf7\xf7\xee\x39\xec\x71\xdc\xc7\x8b\x4e\xa5\xa7\x4c\xb6\xc5\x74\x31\x6a\x23\xa6\xb9\x07\x7f\xd8\x0b\x36\xf3\x8f\x03\x3a\x89\xd9\x9e\x31\xc9\x33\xf2\x0f\xbb\xc7\xc7\x89\xe5\x40\x50\x9b\x10\x00\xb5\x8b\x3b\x58\x98\x80\xff\xa1\xb8\xa4\x10\xc1\xf0\x48\x1c\x0c\x88\xae\x4c\x08\x94\x3f\x90\x68\x5e\x80\xff\xd0\xad\x9f\xb0\xc6\x0a\x74\x0a\x04\x9c\x44\xbf\x90\x04\x42\x27\xd0\x00\xd1\x09\xfa\x19\x71\x2e\xed\x53\xc0\xd1\x8e\x10\xe8\x04\xfc\xfb\xb2\xc3\x63\xa5\x76\x02\x26\x31\x17\x7f\xe8\x56\xe6\xc6\x87\x41\x7c\xe9\x4f\xaf\xbb\xd1\x5a\x24\xaf\xeb\xf5\x3f\x20\xa3\xfe\x99\xb4\x3f\x94\x41\x80\xd4\x5e\x16\x84\xd1\x19\x85\xc7\x29\xd2\x99\x72\xb3\xaa\x52\x5b\xab\x83\xfa\x21\x02\x8c\x17\x14\x23\x0f\x01\x47\x47\x65\xcc\xfb\x84\xf6\x42\x8e\xe9\xa9\x93\x7a\xce\xd1\x93\xff\x5e\x12\x7e\x0f\xe3\xa4\x54\x9f\xbd\xb2\x9d\x0b\xa5\xb6\x65\x8e\xb1\xf2\x36\x2c\x3a\xe6\x7a\x53\x2a\x5e\xfa\xb6\xfa\x31\xf5\xe1\xd6\x44\xc2\x4a\xb0\x22\x9f\xa0\xe4\x1f\x2d\xea\x26\xaf\x62\x5a\x30\x12\x5c\xc4\xfe\x9e\xba\x78\xb5\x58\x26\xda\x5c\x79\xb3\x32\x7e\x30\xc5\x19\x5c\x3d\x36\x0b\x5c\xf9\x45\x98\x13\xab\xba\x7a\xbc\xda\x31\x0f\xa5\x46\x7d\x45\x11\x09\xee\x2a\x25\x7d\x2a\xe2\xbd\x02\x59\xba\x7a\x0b\xa4\xeb\xab\xe7\xc9\x1f\xe0\x12\xeb\xc8\xc0\xcb\x30\x2d\x2f\xce\x30\xfa\xd9\xbb\xe1\xc5\x76\x25\x6c\xd9\xd5\x95\x38\xbb\xc9\x4c\x0b\xde\xa3\x20\x3e\x6d\x5f\x6a\xe0\xaa\x69\xdc\x04\x5a\x2f\x8c\x4e\x7c\x7b\x4e\xeb\x37\xd1\xa9\x10\x03\xfb\x5d\x95\xea\x53\xb1\xd3\xf7\xfd\xed\x00\x30\x61\xa1\x2b\x19\xbb\x64\x95\xc9\x11\x89\xd6\xb6\xf3\x6b\xcc\xcb\x47\x9f\x51\xb2\x47\x0f\xd2\x2e\xb3\x87\x5d\xd1\x9c\x84\xb7\xa6\xc3\xe9\xdb\x19\x8e\x55\xdc\x68\x65\x90\x56\x7b\x39\xd0\xd7\xb1\x39\xc2\xf5\x3a\xc6\x7a\xfe\x39\x5b\x93\xd0\x0b\x07\x10\x8f\xda\x38\xd1\x33\x9c\xa9\x39\x63\x8b\xf5\xb8\x73\x52\xc2\x07\x6a\xf9\xed\x09\xbb\xe7\x71\x65\x24\x75\x1a\x9c\x75\xff\x11\x3c\xf9\xd6\xc8\x22\x23\x6e\x8f\x43\xa2\x66\x21\xfd\xda\x5e\x3e\xf2\xef\xed\xd5\xb1\xa6\x26\x7d\x1a\x54\xd3\x8a\x65\x62\x73\xea\x7c\x73\x60\x7b\x90\x13\x66\x54\x81\x45\x23\x5f\x65\x82\x59\x74\x2f\x4a\xd0\xdf\xc5\xa1\xf2\x48\xc2\x95\x23\x33\x09\x17\x64\x3d\xde\x9a\xf0\x1c\x39\xd1\x63\x52\xdd\x70\xa8\xef\x55\x98\x94\xe5\x2c\xc7\xb9\x7d\xd2\x9a\x0b\x87\xd3\x4e\xc6\x4f\xd9\xd6\xe8\x64\x9c\x1e\xcc\x7c\xfe\x9c\xaf\x15\xfb\x7e\x5f\xce\x9d\x36\xae\x7e\xd5\x85\x47\xec\x54\x3b\xe9\xfc\xd4\xfc\x9e\x6b\x83\xf2\xb9\xf9\xcb\x8b\xd5\xb4\xf8\x99\x7a\x9c\x57\x99\xa9\xc1\xb6\x48\x57\x6e\xd2\xce\x09\x8b\x3e\x79\x4a\xc0\x20\x41\x2d\x40\xe2\xb9\x47\x66\xbd\x54\xe9\xf0\x38\x58\x3b\xc3\x0e\x7c\xdb\x3e\xe1\x81\x9a\x1b\xbd\x96\x0d\x53\x5d\x9f\x74\xc3\xdf\x84\xdb\xb2\x2f\xe7\xf1\x9b\xe3\x59\xc3\x73\x5f\xac\xe3\x66\xed\xae\x68\x19\xf8\xeb\x95\x6d\xe3\x0b\x51\x18\xd8\xb9\x68\xe9\x09\xf5\xbb\x5e\x60\x0b\xfd\xef\xca\xf1\xaf\x00\x00\x00\xff\xff\x9a\x6b\x54\xf1\x8f\x08\x00\x00")

func insecure_pem() ([]byte, error) {
	return bindata_read(
		_insecure_pem,
		"insecure.pem",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() ([]byte, error){
	"insecure-key.pem": insecure_key_pem,
	"insecure.csr":     insecure_csr,
	"insecure.pem":     insecure_pem,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func     func() ([]byte, error)
	Children map[string]*_bintree_t
}

var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"insecure-key.pem": &_bintree_t{insecure_key_pem, map[string]*_bintree_t{}},
	"insecure.csr":     &_bintree_t{insecure_csr, map[string]*_bintree_t{}},
	"insecure.pem":     &_bintree_t{insecure_pem, map[string]*_bintree_t{}},
}}
