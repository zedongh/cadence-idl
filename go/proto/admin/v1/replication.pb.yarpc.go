// The MIT License (MIT)
// 
// Copyright (c) 2021 Uber Technologies, Inc.
// 
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
// 
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
// 
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Code generated by protoc-gen-yarpc-go. DO NOT EDIT.
// source: uber/cadence/admin/v1/replication.proto

package adminv1

var yarpcFileDescriptorClosure90118d56a5f1c507 = [][]byte{
	// uber/cadence/admin/v1/replication.proto
	[]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x58, 0x4b, 0x73, 0x1a, 0xc7,
		0x16, 0xf6, 0x80, 0x78, 0xe8, 0x08, 0x01, 0x6a, 0xc9, 0x16, 0x96, 0xac, 0xba, 0x18, 0xbf, 0x64,
		0xb9, 0x2e, 0x5c, 0xc9, 0xd7, 0xf7, 0x59, 0x29, 0xd7, 0x58, 0xa0, 0x68, 0x62, 0x3d, 0x1b, 0x2c,
		0x47, 0xa9, 0x4a, 0x4d, 0x8d, 0x66, 0x5a, 0x62, 0x4a, 0x30, 0x43, 0x4d, 0x37, 0xc8, 0xec, 0x92,
		0x7d, 0xaa, 0xb2, 0xc9, 0x2e, 0xcb, 0xfc, 0x8d, 0xac, 0xb3, 0x4d, 0x7e, 0x52, 0x6a, 0xba, 0x7b,
		0x80, 0x81, 0x01, 0xcb, 0xe5, 0x45, 0x76, 0xf4, 0x39, 0xdf, 0x79, 0xf4, 0xe9, 0xd3, 0xdf, 0x69,
		0x06, 0x9e, 0x75, 0x2f, 0x88, 0x57, 0x31, 0x0d, 0x8b, 0x38, 0x26, 0xa9, 0x18, 0x56, 0xdb, 0x76,
		0x2a, 0xbd, 0xed, 0x8a, 0x47, 0x3a, 0x2d, 0xdb, 0x34, 0x98, 0xed, 0x3a, 0xe5, 0x8e, 0xe7, 0x32,
		0x17, 0xdd, 0xf5, 0x81, 0x65, 0x09, 0x2c, 0x73, 0x60, 0xb9, 0xb7, 0xbd, 0xf6, 0xb7, 0x2b, 0xd7,
		0xbd, 0x6a, 0x91, 0x0a, 0x07, 0x5d, 0x74, 0x2f, 0x2b, 0xcc, 0x6e, 0x13, 0xca, 0x8c, 0x76, 0x47,
		0xd8, 0xad, 0x15, 0xc3, 0x01, 0x3a, 0xb6, 0xef, 0xde, 0x74, 0xdb, 0xed, 0xc0, 0x73, 0x34, 0xc2,
		0x72, 0xdb, 0x86, 0x1d, 0x20, 0x1e, 0x45, 0x27, 0xd9, 0xb4, 0x29, 0x73, 0xbd, 0xbe, 0x00, 0x95,
		0x7e, 0x8a, 0xc1, 0x32, 0x1e, 0xa6, 0x7d, 0x48, 0x28, 0x35, 0xae, 0x08, 0x45, 0x75, 0x58, 0x1a,
		0xd9, 0x8d, 0xce, 0x0c, 0x7a, 0x4d, 0x0b, 0x4a, 0x31, 0xbe, 0xb9, 0xb0, 0xf3, 0xb4, 0x1c, 0xb9,
		0xa9, 0xf2, 0x88, 0x9b, 0x86, 0x41, 0xaf, 0x71, 0xde, 0x0b, 0x0b, 0x28, 0xfa, 0x2f, 0xdc, 0x6f,
		0x19, 0x94, 0xe9, 0x1e, 0x61, 0x9e, 0x4d, 0x7a, 0xc4, 0xd2, 0xdb, 0x22, 0x9e, 0x6e, 0x5b, 0x85,
		0x58, 0x51, 0xd9, 0x8c, 0xe3, 0x7b, 0x3e, 0x00, 0x07, 0x7a, 0x99, 0x8e, 0x66, 0xa1, 0xfb, 0x90,
		0x6e, 0x1a, 0x54, 0x6f, 0xbb, 0x1e, 0x29, 0xc4, 0x8b, 0xca, 0x66, 0x1a, 0xa7, 0x9a, 0x06, 0x3d,
		0x74, 0x3d, 0x82, 0x30, 0x2c, 0xd1, 0xbe, 0x63, 0xea, 0xb4, 0x69, 0x78, 0x96, 0x4e, 0x99, 0xc1,
		0xba, 0xb4, 0x30, 0x57, 0x54, 0x66, 0xa4, 0x5a, 0xef, 0x3b, 0x66, 0xdd, 0x87, 0xd7, 0x39, 0x1a,
		0xe7, 0x68, 0x58, 0x50, 0xfa, 0x31, 0x09, 0xb9, 0xb1, 0xfd, 0xa0, 0x2f, 0x61, 0xde, 0x2f, 0x83,
		0xce, 0xfa, 0x1d, 0x52, 0x50, 0x8a, 0xca, 0x66, 0x76, 0x67, 0xeb, 0x76, 0xa5, 0x68, 0xf4, 0x3b,
		0x04, 0xa7, 0x99, 0xfc, 0x85, 0x1e, 0x43, 0x96, 0xba, 0x5d, 0xcf, 0x24, 0xbc, 0xac, 0xc3, 0xbd,
		0x67, 0x84, 0xd4, 0xb7, 0xd0, 0x2c, 0xf4, 0x1a, 0x16, 0x4d, 0x8f, 0xc8, 0xf2, 0xdb, 0x6d, 0xb1,
		0xed, 0x85, 0x9d, 0xb5, 0xb2, 0xe8, 0x9d, 0x72, 0xd0, 0x3b, 0xe5, 0x46, 0xd0, 0x3b, 0x38, 0x13,
		0x18, 0xf8, 0x22, 0x64, 0xc2, 0x3d, 0xd1, 0x0f, 0x22, 0x8c, 0xc1, 0x98, 0x67, 0x5f, 0x74, 0x19,
		0x09, 0x8a, 0xf3, 0x62, 0x4a, 0xf2, 0x55, 0x6e, 0xe4, 0x67, 0xa1, 0x0e, 0x4c, 0xf6, 0xef, 0xe0,
		0x15, 0x2b, 0x42, 0x8e, 0xbe, 0x53, 0xe0, 0xe1, 0x44, 0xf5, 0x27, 0x02, 0x26, 0x78, 0xc0, 0x7f,
		0xde, 0xee, 0x34, 0x26, 0x22, 0x6f, 0xd0, 0x59, 0x00, 0xd4, 0x03, 0x0e, 0xd0, 0x0d, 0x93, 0xd9,
		0x3d, 0x9b, 0xf5, 0x27, 0xa2, 0x27, 0x79, 0xf4, 0xed, 0x19, 0xd1, 0x55, 0x69, 0x3a, 0x11, 0x7a,
		0x8d, 0x4e, 0xd5, 0xa2, 0x36, 0xac, 0xc9, 0xbb, 0x24, 0x22, 0xf6, 0x76, 0x46, 0x83, 0xa6, 0x78,
		0xd0, 0xf2, 0x94, 0xa0, 0xfb, 0xc2, 0xd0, 0xf7, 0x78, 0xb6, 0x13, 0x8a, 0xb8, 0xda, 0x8c, 0x56,
		0x21, 0x17, 0xd6, 0x2e, 0x0d, 0xbb, 0xe5, 0xf6, 0x88, 0xa7, 0xb7, 0x0d, 0xef, 0x9a, 0x78, 0xa3,
		0xe1, 0xd2, 0x3c, 0x5c, 0x65, 0x4a, 0xb8, 0x3d, 0x69, 0x78, 0xc8, 0xed, 0x42, 0xf1, 0x0a, 0x97,
		0x53, 0x74, 0x6f, 0x32, 0x00, 0xc3, 0x00, 0xa5, 0x5f, 0x63, 0xb0, 0x12, 0xd5, 0x19, 0xe8, 0x14,
		0xf2, 0xb2, 0xcd, 0xdc, 0x0e, 0xf1, 0x78, 0xfb, 0xc9, 0xdb, 0xf1, 0x74, 0x66, 0x83, 0x1d, 0x07,
		0x68, 0x9c, 0xb3, 0xc2, 0x02, 0x94, 0x85, 0x98, 0xbc, 0x14, 0xf3, 0x38, 0x66, 0x5b, 0xe8, 0x25,
		0x24, 0x05, 0x44, 0xde, 0x81, 0xf5, 0x31, 0xc7, 0x1d, 0x7b, 0xe8, 0x16, 0x4b, 0x28, 0x7a, 0x02,
		0x59, 0xd3, 0x75, 0x2e, 0xed, 0x2b, 0xbd, 0x47, 0x3c, 0xea, 0x67, 0x35, 0xc7, 0x6f, 0xd9, 0xa2,
		0x90, 0x9e, 0x09, 0x21, 0x7a, 0x0e, 0xf9, 0x41, 0x59, 0x03, 0x60, 0x82, 0x03, 0x73, 0x81, 0x3c,
		0x80, 0xfe, 0x0f, 0xee, 0x77, 0x3c, 0xd2, 0xb3, 0xdd, 0x2e, 0xd5, 0x27, 0x6c, 0x92, 0xdc, 0x66,
		0x35, 0x00, 0xec, 0x85, 0x6d, 0x4b, 0x3f, 0x2b, 0xb0, 0x31, 0xb3, 0xcf, 0xfd, 0x7c, 0x25, 0x2b,
		0x98, 0xad, 0x2e, 0x65, 0xc4, 0xe3, 0x55, 0x9c, 0xc7, 0x8b, 0x42, 0xba, 0x2b, 0x84, 0x3e, 0x11,
		0x8a, 0xab, 0x26, 0x2b, 0x94, 0xc0, 0x29, 0xbe, 0xd6, 0x2c, 0xf4, 0x1f, 0x98, 0x1f, 0xcc, 0x91,
		0x5b, 0xb0, 0xc5, 0x10, 0x5c, 0xfa, 0x3d, 0x01, 0x6b, 0xd3, 0xef, 0x01, 0x5a, 0x87, 0x79, 0x79,
		0xc4, 0xb6, 0x25, 0xb3, 0x4a, 0x0b, 0x81, 0x66, 0xa1, 0x77, 0x80, 0x6e, 0x5c, 0xef, 0xfa, 0xb2,
		0xe5, 0xde, 0xe8, 0xe4, 0x03, 0x31, 0xbb, 0xbc, 0x03, 0x62, 0x91, 0xfc, 0x2b, 0x0e, 0xea, 0xbd,
		0x84, 0xd7, 0x02, 0x34, 0x5e, 0xba, 0x19, 0x17, 0xa1, 0x02, 0xa4, 0x82, 0xd2, 0xc6, 0x79, 0x69,
		0x83, 0x25, 0x7a, 0x08, 0x19, 0x6a, 0x36, 0x89, 0xd5, 0x6d, 0x11, 0x5e, 0x05, 0x71, 0xac, 0x0b,
		0x03, 0x99, 0x66, 0x21, 0x15, 0xb2, 0x43, 0x08, 0x27, 0xcf, 0xc4, 0x47, 0xcb, 0xb1, 0x38, 0xb0,
		0xe0, 0xec, 0xb9, 0x01, 0x40, 0x99, 0xe1, 0x31, 0x11, 0x43, 0x9c, 0xee, 0xbc, 0x94, 0x68, 0x16,
		0xfa, 0x02, 0x32, 0x81, 0x9a, 0xfb, 0x4f, 0x7d, 0xd4, 0xff, 0x82, 0xc4, 0x73, 0xef, 0x5f, 0xc1,
		0x32, 0x9f, 0x84, 0x4d, 0x62, 0x78, 0xec, 0x82, 0x18, 0x4c, 0x78, 0x49, 0x7f, 0xd4, 0xcb, 0x92,
		0x6f, 0xb6, 0x1f, 0x58, 0x71, 0x5f, 0xff, 0x82, 0x94, 0x45, 0x98, 0x61, 0xb7, 0x68, 0x61, 0x9e,
		0xdb, 0x3f, 0x88, 0xac, 0xfa, 0x89, 0xd1, 0x6f, 0xb9, 0x86, 0x85, 0x03, 0xb0, 0x5f, 0x61, 0x83,
		0x31, 0xd2, 0xee, 0xb0, 0x02, 0x88, 0x46, 0x92, 0x4b, 0xf4, 0x1a, 0x32, 0x3c, 0x3b, 0xbf, 0xc9,
		0xbb, 0x1e, 0x29, 0x2c, 0xcc, 0x70, 0xbb, 0x27, 0x30, 0x78, 0xc1, 0xb7, 0x90, 0x0b, 0xf4, 0x0f,
		0x58, 0xe1, 0x0e, 0xfc, 0x63, 0x25, 0x9e, 0x6e, 0x5b, 0xc4, 0x61, 0x36, 0xeb, 0x17, 0x32, 0xbc,
		0x77, 0x90, 0xaf, 0x7b, 0xcf, 0x55, 0x9a, 0xd4, 0xa0, 0x23, 0xc8, 0xc9, 0xf3, 0xd5, 0x25, 0x01,
		0x16, 0x16, 0x79, 0xd4, 0x27, 0x53, 0x48, 0x44, 0x5e, 0x2c, 0x49, 0xa4, 0x38, 0xdb, 0x0b, 0xad,
		0x4b, 0xdf, 0xc7, 0x61, 0x75, 0x0a, 0xc9, 0xa2, 0x55, 0x48, 0x05, 0x83, 0x57, 0xe1, 0xe7, 0x9a,
		0x64, 0x62, 0xe4, 0x86, 0xfa, 0x3c, 0x76, 0xab, 0x3e, 0x8f, 0x7f, 0x6e, 0x9f, 0x7f, 0x0b, 0x77,
		0xc7, 0x36, 0xae, 0xdb, 0x8c, 0xb4, 0xfd, 0x21, 0xed, 0x3f, 0xb6, 0x9e, 0xdf, 0x6a, 0xfb, 0x1a,
		0x23, 0x6d, 0xbc, 0xdc, 0x9b, 0x90, 0x51, 0xf4, 0x0a, 0x92, 0xa4, 0x47, 0x1c, 0x16, 0xcc, 0xe0,
		0x8d, 0x68, 0xea, 0x34, 0x98, 0xf1, 0xa6, 0xe5, 0x5e, 0x60, 0x09, 0x46, 0xbb, 0x90, 0x75, 0xc8,
		0x8d, 0xee, 0x75, 0x1d, 0x5d, 0x9a, 0x27, 0x6f, 0x63, 0x9e, 0x71, 0xc8, 0x0d, 0xee, 0x3a, 0x35,
		0x6e, 0x52, 0xfa, 0x45, 0x81, 0xc2, 0xb4, 0xc9, 0x33, 0x9b, 0x53, 0xa2, 0x48, 0x39, 0x16, 0x4d,
		0xca, 0x9f, 0xfb, 0x4c, 0x2a, 0xfd, 0xa0, 0xc0, 0x72, 0x38, 0xcb, 0x86, 0x7b, 0x4d, 0x1c, 0x3f,
		0xc1, 0x80, 0x68, 0xc5, 0xcb, 0x37, 0x81, 0xd3, 0x92, 0x69, 0x29, 0xfa, 0x1a, 0x72, 0x63, 0xc3,
		0x58, 0x32, 0xde, 0xa7, 0x4e, 0x60, 0x9c, 0x0d, 0xcf, 0xdf, 0xd2, 0x6f, 0xe1, 0x07, 0x39, 0x7f,
		0x0c, 0x3a, 0x97, 0xee, 0x5f, 0xc2, 0xc1, 0xeb, 0xa3, 0x2f, 0xde, 0x38, 0xe7, 0x88, 0xe1, 0x2b,
		0x76, 0xe4, 0x16, 0xcd, 0x85, 0x6e, 0xd1, 0x08, 0x73, 0x27, 0xc2, 0xcc, 0xfd, 0x18, 0xb2, 0x97,
		0xb6, 0x47, 0x99, 0xe8, 0xa9, 0x21, 0xaf, 0x66, 0xb8, 0x94, 0x77, 0x8d, 0x66, 0xa1, 0x12, 0x2c,
		0x3a, 0xe4, 0xc3, 0x08, 0x28, 0x25, 0x08, 0xde, 0x17, 0x06, 0x98, 0xf1, 0x19, 0x90, 0x9e, 0x98,
		0x01, 0x7e, 0xf7, 0xe5, 0x47, 0x0b, 0xc9, 0x0f, 0x75, 0x74, 0x7a, 0x2a, 0xe1, 0xe9, 0xf9, 0x19,
		0x7f, 0x4e, 0x02, 0xd3, 0x8e, 0xe7, 0x9a, 0x84, 0xd2, 0xb0, 0x69, 0x7c, 0x68, 0x7a, 0x12, 0xe8,
		0x07, 0xa6, 0xa5, 0xb7, 0x90, 0x1b, 0x7b, 0x16, 0x84, 0xc7, 0xb8, 0xf2, 0x29, 0x63, 0xdc, 0x81,
		0x15, 0x79, 0xf9, 0xab, 0x07, 0xa7, 0xbb, 0x6e, 0xd7, 0x61, 0x35, 0x87, 0x79, 0x7d, 0xb4, 0x02,
		0x09, 0xd3, 0x5f, 0x49, 0xba, 0x13, 0x8b, 0x59, 0x2f, 0x89, 0xc9, 0xb7, 0x48, 0x3c, 0xe2, 0x2d,
		0xb2, 0xf5, 0xc7, 0x64, 0xaf, 0xf2, 0xd6, 0x78, 0x08, 0x1b, 0xb8, 0x76, 0x72, 0xa0, 0xed, 0xaa,
		0x0d, 0xed, 0xf8, 0x48, 0x6f, 0xa8, 0xf5, 0xb7, 0x7a, 0xe3, 0xfc, 0xa4, 0xa6, 0x6b, 0x47, 0x67,
		0xea, 0x81, 0x56, 0xcd, 0xdf, 0x41, 0x45, 0x78, 0x10, 0x0d, 0xa9, 0x1e, 0x1f, 0xaa, 0xda, 0x51,
		0x5e, 0x99, 0xee, 0x64, 0x5f, 0xab, 0x37, 0x8e, 0xf1, 0x79, 0x3e, 0x86, 0x5e, 0xc0, 0xb3, 0x68,
		0x48, 0xfd, 0xfc, 0x68, 0x57, 0xaf, 0xef, 0xab, 0xb8, 0xaa, 0xd7, 0x1b, 0x6a, 0xe3, 0x5d, 0x3d,
		0x1f, 0x47, 0xcf, 0xe0, 0xd1, 0x0c, 0xb0, 0xba, 0xdb, 0xd0, 0xce, 0xb4, 0xc6, 0x79, 0x7e, 0x0e,
		0x6d, 0xc1, 0xd3, 0x99, 0x81, 0xf5, 0xc3, 0x5a, 0x43, 0xad, 0xaa, 0x0d, 0x35, 0x9f, 0x40, 0x8f,
		0xa1, 0x38, 0x1b, 0x7b, 0xb6, 0x93, 0x4f, 0xa2, 0xe7, 0xf0, 0x24, 0x1a, 0xb5, 0xa7, 0x6a, 0x07,
		0xc7, 0x67, 0x35, 0xac, 0x1f, 0xaa, 0xf8, 0x6d, 0x0d, 0xe7, 0x53, 0x5b, 0x36, 0xe4, 0xc6, 0x9e,
		0xc7, 0xe8, 0x01, 0x14, 0x44, 0x51, 0xf4, 0xe3, 0x93, 0x1a, 0x16, 0x2e, 0x86, 0x85, 0x5c, 0x87,
		0xd5, 0x09, 0xed, 0x2e, 0xae, 0xa9, 0x8d, 0x5a, 0x5e, 0x89, 0x54, 0xbe, 0x3b, 0xa9, 0xfa, 0xca,
		0xd8, 0xd6, 0x11, 0xa4, 0xaa, 0x07, 0xa7, 0xfc, 0xc0, 0x56, 0x20, 0x5f, 0x3d, 0x38, 0x1d, 0x3f,
		0xa3, 0x02, 0xac, 0x0c, 0xa4, 0x23, 0xf9, 0xe7, 0x15, 0xb4, 0x0c, 0xb9, 0x81, 0x46, 0x1e, 0x58,
		0xec, 0xcd, 0xbf, 0xbf, 0x79, 0x75, 0x65, 0xb3, 0x66, 0xf7, 0xa2, 0x6c, 0xba, 0xed, 0xca, 0xe8,
		0xc7, 0x87, 0xbf, 0xdb, 0x56, 0xab, 0x72, 0xe5, 0x8a, 0xcf, 0x1d, 0x83, 0x2f, 0x11, 0xff, 0xe7,
		0x3f, 0x7a, 0xdb, 0x17, 0x49, 0x2e, 0x7f, 0xf9, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x40, 0x11,
		0xc6, 0x58, 0x56, 0x11, 0x00, 0x00,
	},
	// google/protobuf/timestamp.proto
	[]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4f, 0xcf, 0xcf, 0x4f,
		0xcf, 0x49, 0xd5, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x2f, 0xc9, 0xcc, 0x4d,
		0x2d, 0x2e, 0x49, 0xcc, 0x2d, 0xd0, 0x03, 0x0b, 0x09, 0xf1, 0x43, 0x14, 0xe8, 0xc1, 0x14, 0x28,
		0x59, 0x73, 0x71, 0x86, 0xc0, 0xd4, 0x08, 0x49, 0x70, 0xb1, 0x17, 0xa7, 0x26, 0xe7, 0xe7, 0xa5,
		0x14, 0x4b, 0x30, 0x2a, 0x30, 0x6a, 0x30, 0x07, 0xc1, 0xb8, 0x42, 0x22, 0x5c, 0xac, 0x79, 0x89,
		0x79, 0xf9, 0xc5, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x10, 0x8e, 0x53, 0x2b, 0x23, 0x97,
		0x70, 0x72, 0x7e, 0xae, 0x1e, 0x9a, 0xa1, 0x4e, 0x7c, 0x70, 0x23, 0x03, 0x40, 0x42, 0x01, 0x8c,
		0x51, 0x46, 0x50, 0x25, 0xe9, 0xf9, 0x39, 0x89, 0x79, 0xe9, 0x7a, 0xf9, 0x45, 0xe9, 0x48, 0x6e,
		0xac, 0x2c, 0x48, 0x2d, 0xd6, 0xcf, 0xce, 0xcb, 0x2f, 0xcf, 0x43, 0xb8, 0xb7, 0x20, 0xe9, 0x07,
		0x23, 0xe3, 0x22, 0x26, 0x66, 0xf7, 0x00, 0xa7, 0x55, 0x4c, 0x72, 0xee, 0x10, 0xdd, 0x01, 0x50,
		0x2d, 0x7a, 0xe1, 0xa9, 0x39, 0x39, 0xde, 0x20, 0x0d, 0x21, 0x20, 0xbd, 0x49, 0x6c, 0x60, 0xb3,
		0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xae, 0x65, 0xce, 0x7d, 0xff, 0x00, 0x00, 0x00,
	},
	// uber/cadence/api/v1/common.proto
	[]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0xdd, 0x6e, 0xdb, 0x36,
		0x14, 0x9e, 0xe2, 0xd8, 0x49, 0x8f, 0xdd, 0xd4, 0x63, 0xd6, 0xd4, 0xc9, 0xfe, 0x3c, 0x03, 0x43,
		0xb3, 0x01, 0x93, 0x10, 0xf7, 0xa6, 0x58, 0x51, 0x0c, 0x4e, 0xec, 0xac, 0x6a, 0xb7, 0xc4, 0x90,
		0x8d, 0x66, 0xdb, 0xc5, 0x04, 0x5a, 0x3c, 0x72, 0x39, 0x4b, 0xa4, 0x40, 0x51, 0x4e, 0x7c, 0xb7,
		0x27, 0xd9, 0xc5, 0x5e, 0x69, 0x2f, 0x34, 0x48, 0xa2, 0x63, 0xa7, 0xf3, 0x90, 0x9b, 0x61, 0x77,
		0xe4, 0xf9, 0x7e, 0xce, 0x47, 0xe1, 0x90, 0x82, 0x76, 0x36, 0x41, 0xe5, 0x04, 0x94, 0xa1, 0x08,
		0xd0, 0xa1, 0x09, 0x77, 0xe6, 0x27, 0x4e, 0x20, 0xe3, 0x58, 0x0a, 0x3b, 0x51, 0x52, 0x4b, 0xb2,
		0x9f, 0x33, 0x6c, 0xc3, 0xb0, 0x69, 0xc2, 0xed, 0xf9, 0xc9, 0xd1, 0x67, 0x53, 0x29, 0xa7, 0x11,
		0x3a, 0x05, 0x65, 0x92, 0x85, 0x0e, 0xcb, 0x14, 0xd5, 0x7c, 0x29, 0xea, 0xbc, 0x81, 0x0f, 0xaf,
		0xa4, 0x9a, 0x85, 0x91, 0xbc, 0x1e, 0xdc, 0x60, 0x90, 0xe5, 0x10, 0xf9, 0x1c, 0xea, 0xd7, 0xa6,
		0xe8, 0x73, 0xd6, 0xb2, 0xda, 0xd6, 0xf1, 0x03, 0x0f, 0x96, 0x25, 0x97, 0x91, 0xc7, 0x50, 0x53,
		0x99, 0xc8, 0xb1, 0xad, 0x02, 0xab, 0xaa, 0x4c, 0xb8, 0xac, 0xd3, 0x81, 0xc6, 0xd2, 0x6c, 0xbc,
		0x48, 0x90, 0x10, 0xd8, 0x16, 0x34, 0x46, 0x63, 0x50, 0xac, 0x73, 0x4e, 0x2f, 0xd0, 0x7c, 0xce,
		0xf5, 0xe2, 0x5f, 0x39, 0x9f, 0xc2, 0xce, 0x90, 0x2e, 0x22, 0x49, 0x59, 0x0e, 0x33, 0xaa, 0x69,
		0x01, 0x37, 0xbc, 0x62, 0xdd, 0x79, 0x01, 0x3b, 0xe7, 0x94, 0x47, 0x99, 0x42, 0x72, 0x00, 0x35,
		0x85, 0x34, 0x95, 0xc2, 0xe8, 0xcd, 0x8e, 0xb4, 0x60, 0x87, 0xa1, 0xa6, 0x3c, 0x4a, 0x8b, 0x84,
		0x0d, 0x6f, 0xb9, 0xed, 0xfc, 0x61, 0xc1, 0xf6, 0x8f, 0x18, 0x4b, 0xf2, 0x12, 0x6a, 0x21, 0xc7,
		0x88, 0xa5, 0x2d, 0xab, 0x5d, 0x39, 0xae, 0x77, 0xbf, 0xb4, 0x37, 0x7c, 0x3f, 0x3b, 0xa7, 0xda,
		0xe7, 0x05, 0x6f, 0x20, 0xb4, 0x5a, 0x78, 0x46, 0x74, 0x74, 0x05, 0xf5, 0xb5, 0x32, 0x69, 0x42,
		0x65, 0x86, 0x0b, 0x93, 0x22, 0x5f, 0x92, 0x2e, 0x54, 0xe7, 0x34, 0xca, 0xb0, 0x08, 0x50, 0xef,
		0x7e, 0xb2, 0xd1, 0xde, 0x1c, 0xd3, 0x2b, 0xa9, 0xdf, 0x6e, 0x3d, 0xb7, 0x3a, 0x7f, 0x5a, 0x50,
		0x7b, 0x85, 0x94, 0xa1, 0x22, 0xdf, 0xbd, 0x17, 0xf1, 0xe9, 0x46, 0x8f, 0x92, 0xfc, 0xff, 0x86,
		0xfc, 0xcb, 0x82, 0xe6, 0x08, 0xa9, 0x0a, 0xde, 0xf5, 0xb4, 0x56, 0x7c, 0x92, 0x69, 0x4c, 0x89,
		0x0f, 0x7b, 0x5c, 0x30, 0xbc, 0x41, 0xe6, 0xdf, 0x89, 0xfd, 0x7c, 0xa3, 0xeb, 0xfb, 0x72, 0xdb,
		0x2d, 0xb5, 0xeb, 0xe7, 0x78, 0xc8, 0xd7, 0x6b, 0x47, 0xbf, 0x02, 0xf9, 0x27, 0xe9, 0x3f, 0x3c,
		0x55, 0x08, 0xbb, 0x7d, 0xaa, 0xe9, 0x69, 0x24, 0x27, 0xe4, 0x1c, 0x1e, 0xa2, 0x08, 0x24, 0xe3,
		0x62, 0xea, 0xeb, 0x45, 0x52, 0x0e, 0xe8, 0x5e, 0xf7, 0x8b, 0x8d, 0x5e, 0x03, 0xc3, 0xcc, 0x27,
		0xda, 0x6b, 0xe0, 0xda, 0xee, 0x76, 0x80, 0xb7, 0xd6, 0x06, 0x78, 0x58, 0x5e, 0x3a, 0x54, 0x6f,
		0x51, 0xa5, 0x5c, 0x0a, 0x57, 0x84, 0x32, 0x27, 0xf2, 0x38, 0x89, 0x96, 0x17, 0x21, 0x5f, 0x93,
		0xa7, 0xf0, 0x28, 0x44, 0xaa, 0x33, 0x85, 0xfe, 0xbc, 0xa4, 0x9a, 0x0b, 0xb7, 0x67, 0xca, 0xc6,
		0xa0, 0xf3, 0x06, 0x9e, 0x8c, 0xb2, 0x24, 0x91, 0x4a, 0x23, 0x3b, 0x8b, 0x38, 0x0a, 0x6d, 0x90,
		0x34, 0xbf, 0xab, 0x53, 0xe9, 0xa7, 0x6c, 0x66, 0x9c, 0xab, 0x53, 0x39, 0x62, 0x33, 0x72, 0x08,
		0xbb, 0xbf, 0xd1, 0x39, 0x2d, 0x80, 0xd2, 0x73, 0x27, 0xdf, 0x8f, 0xd8, 0xac, 0xf3, 0x7b, 0x05,
		0xea, 0x1e, 0x6a, 0xb5, 0x18, 0xca, 0x88, 0x07, 0x0b, 0xd2, 0x87, 0x26, 0x17, 0x5c, 0x73, 0x1a,
		0xf9, 0x5c, 0x68, 0x54, 0x73, 0x5a, 0xa6, 0xac, 0x77, 0x0f, 0xed, 0xf2, 0x79, 0xb1, 0x97, 0xcf,
		0x8b, 0xdd, 0x37, 0xcf, 0x8b, 0xf7, 0xc8, 0x48, 0x5c, 0xa3, 0x20, 0x0e, 0xec, 0x4f, 0x68, 0x30,
		0x93, 0x61, 0xe8, 0x07, 0x12, 0xc3, 0x90, 0x07, 0x79, 0xcc, 0xa2, 0xb7, 0xe5, 0x11, 0x03, 0x9d,
		0xad, 0x90, 0xbc, 0x6d, 0x4c, 0x6f, 0x78, 0x9c, 0xc5, 0xab, 0xb6, 0x95, 0x7b, 0xdb, 0x1a, 0xc9,
		0x6d, 0xdb, 0xaf, 0x56, 0x2e, 0x54, 0x6b, 0x8c, 0x13, 0x9d, 0xb6, 0xb6, 0xdb, 0xd6, 0x71, 0xf5,
		0x96, 0xda, 0x33, 0x65, 0xf2, 0x12, 0x3e, 0x16, 0x52, 0xf8, 0x2a, 0x3f, 0x3a, 0x9d, 0x44, 0xe8,
		0xa3, 0x52, 0x52, 0xf9, 0xe5, 0x93, 0x92, 0xb6, 0xaa, 0xed, 0xca, 0xf1, 0x03, 0xaf, 0x25, 0xa4,
		0xf0, 0x96, 0x8c, 0x41, 0x4e, 0xf0, 0x4a, 0x9c, 0xbc, 0x86, 0x7d, 0xbc, 0x49, 0x78, 0x19, 0x64,
		0x15, 0xb9, 0x76, 0x5f, 0x64, 0xb2, 0x52, 0x2d, 0x53, 0x7f, 0x7d, 0x0d, 0x8d, 0xf5, 0x99, 0x22,
		0x87, 0xf0, 0x78, 0x70, 0x71, 0x76, 0xd9, 0x77, 0x2f, 0xbe, 0xf7, 0xc7, 0x3f, 0x0f, 0x07, 0xbe,
		0x7b, 0xf1, 0xb6, 0xf7, 0x83, 0xdb, 0x6f, 0x7e, 0x40, 0x8e, 0xe0, 0xe0, 0x2e, 0x34, 0x7e, 0xe5,
		0xb9, 0xe7, 0x63, 0xef, 0xaa, 0x69, 0x91, 0x03, 0x20, 0x77, 0xb1, 0xd7, 0xa3, 0xcb, 0x8b, 0xe6,
		0x16, 0x69, 0xc1, 0x47, 0x77, 0xeb, 0x43, 0xef, 0x72, 0x7c, 0xf9, 0xac, 0x59, 0x39, 0xfd, 0x09,
		0x9e, 0x04, 0x32, 0xde, 0x34, 0xe4, 0xa7, 0xbb, 0xbd, 0x84, 0x0f, 0xf3, 0xf4, 0x43, 0xeb, 0x97,
		0x93, 0x29, 0xd7, 0xef, 0xb2, 0x89, 0x1d, 0xc8, 0xd8, 0x59, 0xff, 0x31, 0x7d, 0xc3, 0x59, 0xe4,
		0x4c, 0x65, 0xf9, 0xbb, 0x31, 0x7f, 0xa9, 0x17, 0x34, 0xe1, 0xf3, 0x93, 0x49, 0xad, 0xa8, 0x3d,
		0xfb, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x8c, 0x5b, 0xe4, 0xc9, 0x06, 0x00, 0x00,
	},
	// google/protobuf/duration.proto
	[]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0xcf, 0xcf, 0x4f,
		0xcf, 0x49, 0xd5, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x4f, 0x29, 0x2d, 0x4a,
		0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x03, 0x8b, 0x08, 0xf1, 0x43, 0xe4, 0xf5, 0x60, 0xf2, 0x4a, 0x56,
		0x5c, 0x1c, 0x2e, 0x50, 0x25, 0x42, 0x12, 0x5c, 0xec, 0xc5, 0xa9, 0xc9, 0xf9, 0x79, 0x29, 0xc5,
		0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x30, 0xae, 0x90, 0x08, 0x17, 0x6b, 0x5e, 0x62, 0x5e,
		0x7e, 0xb1, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x6b, 0x10, 0x84, 0xe3, 0xd4, 0xcc, 0xc8, 0x25, 0x9c,
		0x9c, 0x9f, 0xab, 0x87, 0x66, 0xa6, 0x13, 0x2f, 0xcc, 0xc4, 0x00, 0x90, 0x48, 0x00, 0x63, 0x94,
		0x21, 0x54, 0x45, 0x7a, 0x7e, 0x4e, 0x62, 0x5e, 0xba, 0x5e, 0x7e, 0x51, 0x3a, 0xc2, 0x81, 0x25,
		0x95, 0x05, 0xa9, 0xc5, 0xfa, 0xd9, 0x79, 0xf9, 0xe5, 0x79, 0x70, 0xc7, 0x16, 0x24, 0xfd, 0x60,
		0x64, 0x5c, 0xc4, 0xc4, 0xec, 0x1e, 0xe0, 0xb4, 0x8a, 0x49, 0xce, 0x1d, 0xa2, 0x39, 0x00, 0xaa,
		0x43, 0x2f, 0x3c, 0x35, 0x27, 0xc7, 0x1b, 0xa4, 0x3e, 0x04, 0xa4, 0x35, 0x89, 0x0d, 0x6c, 0x94,
		0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xef, 0x8a, 0xb4, 0xc3, 0xfb, 0x00, 0x00, 0x00,
	},
	// uber/cadence/api/v1/domain.proto
	[]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x56, 0xd1, 0x6e, 0xdb, 0x36,
		0x14, 0x9d, 0xec, 0x24, 0x73, 0xae, 0x1c, 0xd7, 0x65, 0x9a, 0x46, 0xf1, 0x86, 0x45, 0x4d, 0x51,
		0xc0, 0x2b, 0x30, 0x79, 0xf1, 0x86, 0xad, 0xdd, 0xb0, 0x07, 0xc7, 0x52, 0x3b, 0x0f, 0x59, 0x16,
		0xc8, 0x6e, 0x30, 0x6c, 0x0f, 0x02, 0x2d, 0xd1, 0x36, 0x51, 0x59, 0x14, 0x28, 0xda, 0x69, 0xde,
		0x86, 0x7d, 0xc4, 0x3e, 0x66, 0x8f, 0xfb, 0xb2, 0x41, 0x14, 0xa5, 0x38, 0xb6, 0x90, 0xf4, 0x8d,
		0xbc, 0xf7, 0x9e, 0xc3, 0xa3, 0xa3, 0x7b, 0x29, 0x81, 0xb9, 0x18, 0x13, 0xde, 0xf1, 0x71, 0x40,
		0x22, 0x9f, 0x74, 0x70, 0x4c, 0x3b, 0xcb, 0xd3, 0x4e, 0xc0, 0xe6, 0x98, 0x46, 0x56, 0xcc, 0x99,
		0x60, 0x68, 0x3f, 0xad, 0xb0, 0x54, 0x85, 0x85, 0x63, 0x6a, 0x2d, 0x4f, 0x5b, 0x5f, 0x4c, 0x19,
		0x9b, 0x86, 0xa4, 0x23, 0x4b, 0xc6, 0x8b, 0x49, 0x27, 0x58, 0x70, 0x2c, 0x28, 0x53, 0xa0, 0xd6,
		0xf1, 0x7a, 0x5e, 0xd0, 0x39, 0x49, 0x04, 0x9e, 0xc7, 0x59, 0xc1, 0xc9, 0x3f, 0x35, 0xd8, 0xb1,
		0xe5, 0x31, 0xa8, 0x01, 0x15, 0x1a, 0x18, 0x9a, 0xa9, 0xb5, 0x77, 0xdd, 0x0a, 0x0d, 0x10, 0x82,
		0xad, 0x08, 0xcf, 0x89, 0x51, 0x91, 0x11, 0xb9, 0x46, 0xaf, 0x61, 0x27, 0x11, 0x58, 0x2c, 0x12,
		0xa3, 0x6a, 0x6a, 0xed, 0x46, 0xf7, 0x99, 0x55, 0xa2, 0xca, 0xca, 0x08, 0x87, 0xb2, 0xd0, 0x55,
		0x00, 0x64, 0x82, 0x1e, 0x90, 0xc4, 0xe7, 0x34, 0x4e, 0xf5, 0x19, 0x5b, 0x92, 0x75, 0x35, 0x84,
		0x8e, 0x41, 0x67, 0xd7, 0x11, 0xe1, 0x1e, 0x99, 0x63, 0x1a, 0x1a, 0xdb, 0xb2, 0x02, 0x64, 0xc8,
		0x49, 0x23, 0xe8, 0x35, 0x6c, 0x05, 0x58, 0x60, 0x63, 0xc7, 0xac, 0xb6, 0xf5, 0xee, 0x8b, 0x7b,
		0xce, 0xb6, 0x6c, 0x2c, 0xb0, 0x13, 0x09, 0x7e, 0xe3, 0x4a, 0x08, 0x9a, 0xc1, 0xf3, 0x6b, 0xc6,
		0xdf, 0x4f, 0x42, 0x76, 0xed, 0x91, 0x0f, 0xc4, 0x5f, 0xa4, 0x27, 0x7a, 0x9c, 0x08, 0x12, 0xc9,
		0x55, 0x4c, 0x38, 0x65, 0x81, 0xf1, 0xa9, 0xa9, 0xb5, 0xf5, 0xee, 0x91, 0x95, 0xd9, 0x66, 0xe5,
		0xb6, 0x59, 0xb6, 0xb2, 0xd5, 0x35, 0x73, 0x16, 0x27, 0x27, 0x71, 0x73, 0x8e, 0x4b, 0x49, 0x81,
		0xfa, 0x50, 0x1f, 0xe3, 0xc0, 0x1b, 0xd3, 0x08, 0x73, 0x4a, 0x12, 0xa3, 0x26, 0x29, 0xcd, 0x52,
		0xb1, 0x67, 0x38, 0x38, 0x53, 0x75, 0xae, 0x3e, 0xbe, 0xdd, 0xa0, 0x3f, 0xe1, 0x70, 0x46, 0x13,
		0xc1, 0xf8, 0x8d, 0x87, 0xb9, 0x3f, 0xa3, 0x4b, 0x1c, 0x7a, 0xca, 0xf8, 0x5d, 0x69, 0xfc, 0xf3,
		0x52, 0xbe, 0x9e, 0xaa, 0x55, 0xd6, 0x1f, 0x28, 0x8e, 0xbb, 0x61, 0xf4, 0x35, 0x3c, 0xd9, 0x20,
		0x5f, 0x70, 0x6a, 0x80, 0x34, 0x1c, 0xad, 0x81, 0xde, 0x71, 0x8a, 0x30, 0xb4, 0x96, 0x34, 0xa1,
		0x63, 0x1a, 0x52, 0xb1, 0xa9, 0x48, 0xff, 0x78, 0x45, 0xc6, 0x2d, 0xcd, 0x9a, 0xa8, 0xef, 0xe0,
		0xb0, 0xec, 0x88, 0x54, 0x57, 0x5d, 0xea, 0x3a, 0xd8, 0x84, 0xa6, 0xd2, 0x2c, 0xd8, 0xc7, 0xbe,
		0xa0, 0x4b, 0xe2, 0xf9, 0xe1, 0x22, 0x11, 0x84, 0x7b, 0xb2, 0x69, 0xf7, 0x24, 0xe6, 0x71, 0x96,
		0xea, 0x67, 0x99, 0x8b, 0xb4, 0x83, 0x2f, 0xa1, 0xa6, 0x0a, 0x13, 0xa3, 0x21, 0xfb, 0xe8, 0xdb,
		0x52, 0xe1, 0x0a, 0xe3, 0x92, 0x38, 0xa4, 0xbe, 0x7c, 0xf7, 0x7d, 0x16, 0x4d, 0xe8, 0x34, 0x6f,
		0x84, 0x82, 0x05, 0x7d, 0x09, 0xcd, 0x09, 0xa6, 0x21, 0x5b, 0x12, 0xee, 0x2d, 0x09, 0x4f, 0xd2,
		0xee, 0x7e, 0x64, 0x6a, 0xed, 0xaa, 0xfb, 0x28, 0x8f, 0x5f, 0x65, 0x61, 0xd4, 0x86, 0x26, 0x4d,
		0xbc, 0x69, 0xc8, 0xc6, 0x38, 0xf4, 0xb2, 0xe9, 0x36, 0x9a, 0xa6, 0xd6, 0xae, 0xb9, 0x0d, 0x9a,
		0xbc, 0x95, 0x61, 0x35, 0x8c, 0x6f, 0x60, 0xaf, 0x20, 0xa5, 0xd1, 0x84, 0x19, 0x8f, 0x65, 0x1b,
		0x95, 0xcf, 0xdb, 0x1b, 0x55, 0x39, 0x88, 0x26, 0xcc, 0xad, 0x4f, 0x56, 0x76, 0xad, 0xef, 0x61,
		0xb7, 0x18, 0x05, 0xd4, 0x84, 0xea, 0x7b, 0x72, 0xa3, 0x46, 0x3c, 0x5d, 0xa2, 0x27, 0xb0, 0xbd,
		0xc4, 0xe1, 0x22, 0x1f, 0xf2, 0x6c, 0xf3, 0x43, 0xe5, 0x95, 0x76, 0x62, 0xc3, 0xf1, 0x03, 0x16,
		0xa0, 0x67, 0x50, 0xbf, 0xe3, 0x79, 0xc6, 0xab, 0xfb, 0xb7, 0x6e, 0x9f, 0xfc, 0xab, 0x81, 0xbe,
		0xd2, 0xe4, 0xe8, 0x17, 0xa8, 0x15, 0x83, 0xa1, 0x49, 0xf7, 0xad, 0x87, 0x06, 0xc3, 0xca, 0x17,
		0xd9, 0x38, 0x17, 0xf8, 0x96, 0x07, 0x7b, 0x77, 0x52, 0x25, 0x8f, 0xf7, 0x6a, 0xf5, 0xf1, 0xf4,
		0xee, 0xc9, 0xbd, 0x67, 0xdd, 0x48, 0xfb, 0x56, 0x2c, 0xf8, 0x5b, 0x83, 0xbd, 0x3b, 0x49, 0xf4,
		0x14, 0x76, 0x38, 0xc1, 0x09, 0x8b, 0xd4, 0x21, 0x6a, 0x87, 0x5a, 0x50, 0x63, 0x31, 0xe1, 0x58,
		0x30, 0xae, 0x9c, 0x2c, 0xf6, 0xe8, 0x27, 0xa8, 0xfb, 0x9c, 0x60, 0x41, 0x02, 0x2f, 0xbd, 0x7c,
		0xe5, 0xc5, 0xa9, 0x77, 0x5b, 0x1b, 0x57, 0xcc, 0x28, 0xbf, 0x99, 0x5d, 0x5d, 0xd5, 0xa7, 0x91,
		0x93, 0xff, 0x2a, 0x50, 0x5f, 0x7d, 0xbf, 0xa5, 0xed, 0xa6, 0x95, 0xb7, 0xdb, 0x08, 0x8c, 0xa2,
		0x34, 0x11, 0x98, 0x0b, 0xaf, 0xb8, 0xfe, 0x95, 0x23, 0xf7, 0xc9, 0x78, 0x9a, 0x63, 0x87, 0x29,
		0xb4, 0x88, 0xa3, 0x2b, 0x38, 0x2a, 0x58, 0xc9, 0x87, 0x98, 0x72, 0xb2, 0x42, 0xfb, 0xf0, 0xd3,
		0x1d, 0xe6, 0x60, 0x47, 0x62, 0x6f, 0x79, 0xbb, 0x70, 0xe0, 0xb3, 0x79, 0x1c, 0x92, 0xd4, 0xaa,
		0x64, 0x86, 0x79, 0xe0, 0xf9, 0x6c, 0x11, 0x09, 0xf9, 0xa9, 0xd8, 0x76, 0xf7, 0x8b, 0xe4, 0x30,
		0xcd, 0xf5, 0xd3, 0x14, 0x7a, 0x01, 0x8d, 0x98, 0x44, 0x01, 0x8d, 0xa6, 0x19, 0x22, 0x31, 0xb6,
		0xcd, 0x6a, 0x7b, 0xdb, 0xdd, 0x53, 0x51, 0x59, 0x9a, 0xbc, 0xfc, 0x4b, 0x83, 0xfa, 0xea, 0x47,
		0x09, 0x1d, 0xc1, 0x81, 0xfd, 0xdb, 0xaf, 0xbd, 0xc1, 0x85, 0x37, 0x1c, 0xf5, 0x46, 0xef, 0x86,
		0xde, 0xe0, 0xe2, 0xaa, 0x77, 0x3e, 0xb0, 0x9b, 0x9f, 0xa0, 0xcf, 0xc1, 0xb8, 0x9b, 0x72, 0x9d,
		0xb7, 0x83, 0xe1, 0xc8, 0x71, 0x1d, 0xbb, 0xa9, 0x6d, 0x66, 0x6d, 0xe7, 0xd2, 0x75, 0xfa, 0xbd,
		0x91, 0x63, 0x37, 0x2b, 0x9b, 0xb4, 0xb6, 0x73, 0xee, 0xa4, 0xa9, 0xea, 0xcb, 0x19, 0x34, 0xd6,
		0x6e, 0xbc, 0xcf, 0xe0, 0xb0, 0xe7, 0xf6, 0x7f, 0x1e, 0x5c, 0xf5, 0xce, 0x4b, 0x55, 0xac, 0x27,
		0xed, 0xc1, 0xb0, 0x77, 0x76, 0x2e, 0x55, 0x94, 0x40, 0x9d, 0x8b, 0x2c, 0x59, 0x39, 0xfb, 0x1d,
		0x0e, 0x7d, 0x36, 0x2f, 0x6b, 0xf5, 0xb3, 0x5a, 0x2f, 0xa6, 0x97, 0xe9, 0x2b, 0xb9, 0xd4, 0xfe,
		0x38, 0x9d, 0x52, 0x31, 0x5b, 0x8c, 0x2d, 0x9f, 0xcd, 0x3b, 0xab, 0x3f, 0x1f, 0x5f, 0xd1, 0x20,
		0xec, 0x4c, 0x59, 0xf6, 0xcb, 0xa0, 0xfe, 0x44, 0x7e, 0xc4, 0x31, 0x5d, 0x9e, 0x8e, 0x77, 0x64,
		0xec, 0x9b, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0x98, 0xc3, 0x6b, 0x64, 0xad, 0x08, 0x00, 0x00,
	},
	// uber/cadence/admin/v1/history.proto
	[]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2e, 0x4d, 0x4a, 0x2d,
		0xd2, 0x4f, 0x4e, 0x4c, 0x49, 0xcd, 0x4b, 0x4e, 0xd5, 0x4f, 0x4c, 0xc9, 0xcd, 0xcc, 0xd3, 0x2f,
		0x33, 0xd4, 0xcf, 0xc8, 0x2c, 0x2e, 0xc9, 0x2f, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
		0x12, 0x05, 0x29, 0xd2, 0x83, 0x2a, 0xd2, 0x03, 0x2b, 0xd2, 0x2b, 0x33, 0x54, 0xf2, 0xe4, 0x12,
		0x0a, 0x4b, 0x2d, 0x2a, 0xce, 0xcc, 0xcf, 0xf3, 0x80, 0x28, 0xf7, 0x2c, 0x49, 0xcd, 0x15, 0x92,
		0xe4, 0xe2, 0x48, 0x2d, 0x4b, 0xcd, 0x2b, 0x89, 0xcf, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60,
		0x0e, 0x62, 0x07, 0xf3, 0x3d, 0x53, 0x84, 0x24, 0xb8, 0xd8, 0xcb, 0x20, 0x1a, 0x24, 0x98, 0x20,
		0x32, 0x50, 0xae, 0x52, 0x09, 0x17, 0x1f, 0xaa, 0x51, 0x42, 0x8a, 0x5c, 0x3c, 0x49, 0x45, 0x89,
		0x79, 0xc9, 0x19, 0xf1, 0x25, 0xf9, 0xd9, 0xa9, 0x79, 0x60, 0xa3, 0x78, 0x82, 0xb8, 0x21, 0x62,
		0x21, 0x20, 0x21, 0x21, 0x7b, 0x2e, 0xd6, 0xcc, 0x92, 0xd4, 0xdc, 0x62, 0x09, 0x26, 0x05, 0x66,
		0x0d, 0x6e, 0x23, 0x4d, 0x3d, 0xac, 0xce, 0xd4, 0xc3, 0x74, 0x63, 0x10, 0x44, 0x9f, 0x93, 0x79,
		0x94, 0x69, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x3e, 0x72, 0x48, 0xe8,
		0x66, 0xa6, 0xe4, 0xe8, 0xa7, 0xe7, 0xeb, 0x83, 0xfd, 0x0f, 0x0f, 0x16, 0x6b, 0x30, 0xa3, 0xcc,
		0x30, 0x89, 0x0d, 0x2c, 0x6e, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x44, 0x14, 0xd7, 0xd4, 0x3e,
		0x01, 0x00, 0x00,
	},
}