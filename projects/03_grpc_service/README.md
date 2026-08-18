# 🚀 gRPC & Protocol Buffers Microservice in Go

## 📝 Overview (সংক্ষিপ্ত পরিচয়)
**gRPC (Google Remote Procedure Call)** হলো গুগল কর্তৃক তৈরি উচ্চ-পারফরম্যান্স ইন্টার-সার্ভিস মাইক্রোসার্ভিস যোগাযোগ ফ্রেমওয়ার্ক। এটি সনাতন REST HTTP/1.1 JSON-এর পরিবর্তে **HTTP/2** এবং **Protocol Buffers (ProtoBuf)** ব্যবহার করে **১০ গুণ দ্রুততর** গতি প্রদান করে।

---

## 🆚 REST API vs gRPC Comparison (পার্থক্য)

| Feature | REST API | gRPC |
| :--- | :--- | :--- |
| **Protocol** | HTTP/1.1 (Text JSON) | HTTP/2 (Binary ProtoBuf) |
| **Data Format** | JSON / XML (Heavy) | Protocol Buffers (Ultra Light) |
| **Speed / Latency** | Standard (~50ms) | Ultra Fast (~5ms — 10x faster) |
| **Streaming Support** | Unidirectional | Bi-directional Streaming |
| **Contract / Schema** | OpenAPI / Swagger (Optional) | Strictly Typed `.proto` Schema |

---

## 🎨 System Architecture Diagram (আর্কিটেকচার চিত্র)

```text
  ┌─────────────────┐                                  ┌─────────────────┐
  │                 │  1. Unary / Streaming RPC        │                 │
  │   gRPC Client   ├─────────────────────────────────►│   gRPC Server   │
  │ (Go / Microsvc) │                                  │ (Go Microservice│
  │                 │  2. Binary ProtoBuf over HTTP/2  │                 │
  │                 │◄─────────────────────────────────┤                 │
  └─────────────────┘                                  └─────────────────┘
```

---

## 📜 Protocol Buffer Definition (`proto/user.proto`)

```protobuf
syntax = "proto3";

package user;

service UserService {
  rpc GetUser (UserRequest) returns (UserResponse);
  rpc CreateUser (CreateUserRequest) returns (UserResponse);
}

message UserRequest {
  int32 id = 1;
}

message UserResponse {
  int32 id = 1;
  string username = 2;
  string email = 3;
}
```

---

## 🧪 How to Run (কোড চালনার নিয়ম)

### 1. Start gRPC Server (সার্ভার চালু করা):
```bash
go run ./projects/03_grpc_service/server.go
```

### 2. Run gRPC Client in another terminal (ক্লায়েন্ট টেস্ট করা):
```bash
go run ./projects/03_grpc_service/client.go
```
