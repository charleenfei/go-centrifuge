// Code generated by protoc-gen-go. DO NOT EDIT.
// source: funding/funding.proto

package funpb

import (
	context "context"
	fmt "fmt"
	math "math"

	document "github.com/centrifuge/go-centrifuge/protobufs/gen/go/document"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type FundingCreatePayload struct {
	Identifier           string       `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Data                 *FundingData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *FundingCreatePayload) Reset()         { *m = FundingCreatePayload{} }
func (m *FundingCreatePayload) String() string { return proto.CompactTextString(m) }
func (*FundingCreatePayload) ProtoMessage()    {}
func (*FundingCreatePayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_e74e73a4c22d632c, []int{0}
}

func (m *FundingCreatePayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FundingCreatePayload.Unmarshal(m, b)
}
func (m *FundingCreatePayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FundingCreatePayload.Marshal(b, m, deterministic)
}
func (m *FundingCreatePayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FundingCreatePayload.Merge(m, src)
}
func (m *FundingCreatePayload) XXX_Size() int {
	return xxx_messageInfo_FundingCreatePayload.Size(m)
}
func (m *FundingCreatePayload) XXX_DiscardUnknown() {
	xxx_messageInfo_FundingCreatePayload.DiscardUnknown(m)
}

var xxx_messageInfo_FundingCreatePayload proto.InternalMessageInfo

func (m *FundingCreatePayload) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *FundingCreatePayload) GetData() *FundingData {
	if m != nil {
		return m.Data
	}
	return nil
}

type FundingUpdatePayload struct {
	Identifier           string       `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	AgreementId          string       `protobuf:"bytes,2,opt,name=agreement_id,json=agreementId,proto3" json:"agreement_id,omitempty"`
	Data                 *FundingData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *FundingUpdatePayload) Reset()         { *m = FundingUpdatePayload{} }
func (m *FundingUpdatePayload) String() string { return proto.CompactTextString(m) }
func (*FundingUpdatePayload) ProtoMessage()    {}
func (*FundingUpdatePayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_e74e73a4c22d632c, []int{1}
}

func (m *FundingUpdatePayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FundingUpdatePayload.Unmarshal(m, b)
}
func (m *FundingUpdatePayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FundingUpdatePayload.Marshal(b, m, deterministic)
}
func (m *FundingUpdatePayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FundingUpdatePayload.Merge(m, src)
}
func (m *FundingUpdatePayload) XXX_Size() int {
	return xxx_messageInfo_FundingUpdatePayload.Size(m)
}
func (m *FundingUpdatePayload) XXX_DiscardUnknown() {
	xxx_messageInfo_FundingUpdatePayload.DiscardUnknown(m)
}

var xxx_messageInfo_FundingUpdatePayload proto.InternalMessageInfo

func (m *FundingUpdatePayload) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *FundingUpdatePayload) GetAgreementId() string {
	if m != nil {
		return m.AgreementId
	}
	return ""
}

func (m *FundingUpdatePayload) GetData() *FundingData {
	if m != nil {
		return m.Data
	}
	return nil
}

type FundingResponse struct {
	Header               *document.ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Data                 *FundingResponseData     `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *FundingResponse) Reset()         { *m = FundingResponse{} }
func (m *FundingResponse) String() string { return proto.CompactTextString(m) }
func (*FundingResponse) ProtoMessage()    {}
func (*FundingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e74e73a4c22d632c, []int{2}
}

func (m *FundingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FundingResponse.Unmarshal(m, b)
}
func (m *FundingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FundingResponse.Marshal(b, m, deterministic)
}
func (m *FundingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FundingResponse.Merge(m, src)
}
func (m *FundingResponse) XXX_Size() int {
	return xxx_messageInfo_FundingResponse.Size(m)
}
func (m *FundingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FundingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FundingResponse proto.InternalMessageInfo

func (m *FundingResponse) GetHeader() *document.ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *FundingResponse) GetData() *FundingResponseData {
	if m != nil {
		return m.Data
	}
	return nil
}

type FundingListResponse struct {
	Header               *document.ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Data                 []*FundingResponseData   `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *FundingListResponse) Reset()         { *m = FundingListResponse{} }
func (m *FundingListResponse) String() string { return proto.CompactTextString(m) }
func (*FundingListResponse) ProtoMessage()    {}
func (*FundingListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e74e73a4c22d632c, []int{3}
}

func (m *FundingListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FundingListResponse.Unmarshal(m, b)
}
func (m *FundingListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FundingListResponse.Marshal(b, m, deterministic)
}
func (m *FundingListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FundingListResponse.Merge(m, src)
}
func (m *FundingListResponse) XXX_Size() int {
	return xxx_messageInfo_FundingListResponse.Size(m)
}
func (m *FundingListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FundingListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FundingListResponse proto.InternalMessageInfo

func (m *FundingListResponse) GetHeader() *document.ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *FundingListResponse) GetData() []*FundingResponseData {
	if m != nil {
		return m.Data
	}
	return nil
}

type Request struct {
	Identifier           string   `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	AgreementId          string   `protobuf:"bytes,2,opt,name=agreement_id,json=agreementId,proto3" json:"agreement_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_e74e73a4c22d632c, []int{4}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *Request) GetAgreementId() string {
	if m != nil {
		return m.AgreementId
	}
	return ""
}

type GetVersionRequest struct {
	Identifier           string   `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	AgreementId          string   `protobuf:"bytes,3,opt,name=agreement_id,json=agreementId,proto3" json:"agreement_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetVersionRequest) Reset()         { *m = GetVersionRequest{} }
func (m *GetVersionRequest) String() string { return proto.CompactTextString(m) }
func (*GetVersionRequest) ProtoMessage()    {}
func (*GetVersionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e74e73a4c22d632c, []int{5}
}

func (m *GetVersionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetVersionRequest.Unmarshal(m, b)
}
func (m *GetVersionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetVersionRequest.Marshal(b, m, deterministic)
}
func (m *GetVersionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetVersionRequest.Merge(m, src)
}
func (m *GetVersionRequest) XXX_Size() int {
	return xxx_messageInfo_GetVersionRequest.Size(m)
}
func (m *GetVersionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetVersionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetVersionRequest proto.InternalMessageInfo

func (m *GetVersionRequest) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *GetVersionRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *GetVersionRequest) GetAgreementId() string {
	if m != nil {
		return m.AgreementId
	}
	return ""
}

type GetListRequest struct {
	Identifier           string   `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetListRequest) Reset()         { *m = GetListRequest{} }
func (m *GetListRequest) String() string { return proto.CompactTextString(m) }
func (*GetListRequest) ProtoMessage()    {}
func (*GetListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e74e73a4c22d632c, []int{6}
}

func (m *GetListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetListRequest.Unmarshal(m, b)
}
func (m *GetListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetListRequest.Marshal(b, m, deterministic)
}
func (m *GetListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetListRequest.Merge(m, src)
}
func (m *GetListRequest) XXX_Size() int {
	return xxx_messageInfo_GetListRequest.Size(m)
}
func (m *GetListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetListRequest proto.InternalMessageInfo

func (m *GetListRequest) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

type GetListVersionRequest struct {
	Identifier           string   `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetListVersionRequest) Reset()         { *m = GetListVersionRequest{} }
func (m *GetListVersionRequest) String() string { return proto.CompactTextString(m) }
func (*GetListVersionRequest) ProtoMessage()    {}
func (*GetListVersionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e74e73a4c22d632c, []int{7}
}

func (m *GetListVersionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetListVersionRequest.Unmarshal(m, b)
}
func (m *GetListVersionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetListVersionRequest.Marshal(b, m, deterministic)
}
func (m *GetListVersionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetListVersionRequest.Merge(m, src)
}
func (m *GetListVersionRequest) XXX_Size() int {
	return xxx_messageInfo_GetListVersionRequest.Size(m)
}
func (m *GetListVersionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetListVersionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetListVersionRequest proto.InternalMessageInfo

func (m *GetListVersionRequest) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *GetListVersionRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

// FundingData is the default funding extension schema
type FundingData struct {
	AgreementId           string   `protobuf:"bytes,1,opt,name=agreement_id,json=agreementId,proto3" json:"agreement_id,omitempty"`
	Amount                string   `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Apr                   string   `protobuf:"bytes,3,opt,name=apr,proto3" json:"apr,omitempty"`
	Days                  string   `protobuf:"bytes,4,opt,name=days,proto3" json:"days,omitempty"`
	Fee                   string   `protobuf:"bytes,5,opt,name=fee,proto3" json:"fee,omitempty"`
	RepaymentDueDate      string   `protobuf:"bytes,7,opt,name=repayment_due_date,json=repaymentDueDate,proto3" json:"repayment_due_date,omitempty"`
	RepaymentOccurredDate string   `protobuf:"bytes,8,opt,name=repayment_occurred_date,json=repaymentOccurredDate,proto3" json:"repayment_occurred_date,omitempty"`
	RepaymentAmount       string   `protobuf:"bytes,9,opt,name=repayment_amount,json=repaymentAmount,proto3" json:"repayment_amount,omitempty"`
	Currency              string   `protobuf:"bytes,10,opt,name=currency,proto3" json:"currency,omitempty"`
	NftAddress            string   `protobuf:"bytes,11,opt,name=nft_address,json=nftAddress,proto3" json:"nft_address,omitempty"`
	PaymentDetailsId      string   `protobuf:"bytes,12,opt,name=payment_details_id,json=paymentDetailsId,proto3" json:"payment_details_id,omitempty"`
	FunderId              string   `protobuf:"bytes,13,opt,name=funder_id,json=funderId,proto3" json:"funder_id,omitempty"`
	BorrowerId            string   `protobuf:"bytes,14,opt,name=borrower_id,json=borrowerId,proto3" json:"borrower_id,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *FundingData) Reset()         { *m = FundingData{} }
func (m *FundingData) String() string { return proto.CompactTextString(m) }
func (*FundingData) ProtoMessage()    {}
func (*FundingData) Descriptor() ([]byte, []int) {
	return fileDescriptor_e74e73a4c22d632c, []int{8}
}

func (m *FundingData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FundingData.Unmarshal(m, b)
}
func (m *FundingData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FundingData.Marshal(b, m, deterministic)
}
func (m *FundingData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FundingData.Merge(m, src)
}
func (m *FundingData) XXX_Size() int {
	return xxx_messageInfo_FundingData.Size(m)
}
func (m *FundingData) XXX_DiscardUnknown() {
	xxx_messageInfo_FundingData.DiscardUnknown(m)
}

var xxx_messageInfo_FundingData proto.InternalMessageInfo

func (m *FundingData) GetAgreementId() string {
	if m != nil {
		return m.AgreementId
	}
	return ""
}

func (m *FundingData) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func (m *FundingData) GetApr() string {
	if m != nil {
		return m.Apr
	}
	return ""
}

func (m *FundingData) GetDays() string {
	if m != nil {
		return m.Days
	}
	return ""
}

func (m *FundingData) GetFee() string {
	if m != nil {
		return m.Fee
	}
	return ""
}

func (m *FundingData) GetRepaymentDueDate() string {
	if m != nil {
		return m.RepaymentDueDate
	}
	return ""
}

func (m *FundingData) GetRepaymentOccurredDate() string {
	if m != nil {
		return m.RepaymentOccurredDate
	}
	return ""
}

func (m *FundingData) GetRepaymentAmount() string {
	if m != nil {
		return m.RepaymentAmount
	}
	return ""
}

func (m *FundingData) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *FundingData) GetNftAddress() string {
	if m != nil {
		return m.NftAddress
	}
	return ""
}

func (m *FundingData) GetPaymentDetailsId() string {
	if m != nil {
		return m.PaymentDetailsId
	}
	return ""
}

func (m *FundingData) GetFunderId() string {
	if m != nil {
		return m.FunderId
	}
	return ""
}

func (m *FundingData) GetBorrowerId() string {
	if m != nil {
		return m.BorrowerId
	}
	return ""
}

type FundingResponseData struct {
	Funding              *FundingData        `protobuf:"bytes,1,opt,name=funding,proto3" json:"funding,omitempty"`
	Signatures           []*FundingSignature `protobuf:"bytes,2,rep,name=signatures,proto3" json:"signatures,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *FundingResponseData) Reset()         { *m = FundingResponseData{} }
func (m *FundingResponseData) String() string { return proto.CompactTextString(m) }
func (*FundingResponseData) ProtoMessage()    {}
func (*FundingResponseData) Descriptor() ([]byte, []int) {
	return fileDescriptor_e74e73a4c22d632c, []int{9}
}

func (m *FundingResponseData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FundingResponseData.Unmarshal(m, b)
}
func (m *FundingResponseData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FundingResponseData.Marshal(b, m, deterministic)
}
func (m *FundingResponseData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FundingResponseData.Merge(m, src)
}
func (m *FundingResponseData) XXX_Size() int {
	return xxx_messageInfo_FundingResponseData.Size(m)
}
func (m *FundingResponseData) XXX_DiscardUnknown() {
	xxx_messageInfo_FundingResponseData.DiscardUnknown(m)
}

var xxx_messageInfo_FundingResponseData proto.InternalMessageInfo

func (m *FundingResponseData) GetFunding() *FundingData {
	if m != nil {
		return m.Funding
	}
	return nil
}

func (m *FundingResponseData) GetSignatures() []*FundingSignature {
	if m != nil {
		return m.Signatures
	}
	return nil
}

type FundingSignature struct {
	Valid                string   `protobuf:"bytes,1,opt,name=valid,proto3" json:"valid,omitempty"`
	OutdatedSignature    string   `protobuf:"bytes,2,opt,name=outdated_signature,json=outdatedSignature,proto3" json:"outdated_signature,omitempty"`
	Identity             string   `protobuf:"bytes,3,opt,name=identity,proto3" json:"identity,omitempty"`
	SignedVersion        string   `protobuf:"bytes,4,opt,name=signed_version,json=signedVersion,proto3" json:"signed_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FundingSignature) Reset()         { *m = FundingSignature{} }
func (m *FundingSignature) String() string { return proto.CompactTextString(m) }
func (*FundingSignature) ProtoMessage()    {}
func (*FundingSignature) Descriptor() ([]byte, []int) {
	return fileDescriptor_e74e73a4c22d632c, []int{10}
}

func (m *FundingSignature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FundingSignature.Unmarshal(m, b)
}
func (m *FundingSignature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FundingSignature.Marshal(b, m, deterministic)
}
func (m *FundingSignature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FundingSignature.Merge(m, src)
}
func (m *FundingSignature) XXX_Size() int {
	return xxx_messageInfo_FundingSignature.Size(m)
}
func (m *FundingSignature) XXX_DiscardUnknown() {
	xxx_messageInfo_FundingSignature.DiscardUnknown(m)
}

var xxx_messageInfo_FundingSignature proto.InternalMessageInfo

func (m *FundingSignature) GetValid() string {
	if m != nil {
		return m.Valid
	}
	return ""
}

func (m *FundingSignature) GetOutdatedSignature() string {
	if m != nil {
		return m.OutdatedSignature
	}
	return ""
}

func (m *FundingSignature) GetIdentity() string {
	if m != nil {
		return m.Identity
	}
	return ""
}

func (m *FundingSignature) GetSignedVersion() string {
	if m != nil {
		return m.SignedVersion
	}
	return ""
}

func init() {
	proto.RegisterType((*FundingCreatePayload)(nil), "fun.FundingCreatePayload")
	proto.RegisterType((*FundingUpdatePayload)(nil), "fun.FundingUpdatePayload")
	proto.RegisterType((*FundingResponse)(nil), "fun.FundingResponse")
	proto.RegisterType((*FundingListResponse)(nil), "fun.FundingListResponse")
	proto.RegisterType((*Request)(nil), "fun.Request")
	proto.RegisterType((*GetVersionRequest)(nil), "fun.GetVersionRequest")
	proto.RegisterType((*GetListRequest)(nil), "fun.GetListRequest")
	proto.RegisterType((*GetListVersionRequest)(nil), "fun.GetListVersionRequest")
	proto.RegisterType((*FundingData)(nil), "fun.FundingData")
	proto.RegisterType((*FundingResponseData)(nil), "fun.FundingResponseData")
	proto.RegisterType((*FundingSignature)(nil), "fun.FundingSignature")
}

func init() { proto.RegisterFile("funding/funding.proto", fileDescriptor_e74e73a4c22d632c) }

var fileDescriptor_e74e73a4c22d632c = []byte{
	// 971 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x96, 0x5f, 0x6f, 0x1b, 0x45,
	0x10, 0xc0, 0x75, 0x75, 0x12, 0x27, 0xe3, 0xd4, 0x4d, 0xb7, 0x49, 0x38, 0x8e, 0xaa, 0x1c, 0x27,
	0x22, 0xda, 0xe0, 0xf8, 0xd2, 0x50, 0x90, 0x0a, 0x4f, 0x57, 0x2a, 0x42, 0xa4, 0x48, 0x04, 0x23,
	0x78, 0x40, 0x95, 0xac, 0x8d, 0x77, 0xcf, 0x39, 0xc9, 0xd9, 0xbd, 0xde, 0xee, 0xa5, 0x58, 0x51,
	0x85, 0xc4, 0x0b, 0xef, 0xae, 0xe0, 0x81, 0x57, 0x3e, 0x00, 0xef, 0x7c, 0x0d, 0xbe, 0x02, 0x6f,
	0x7c, 0x09, 0xb4, 0x7f, 0xee, 0x7c, 0x8e, 0xe3, 0xc6, 0x91, 0xf2, 0x64, 0xef, 0xcc, 0xec, 0xcc,
	0x6f, 0x67, 0x67, 0x76, 0x0e, 0x36, 0xe2, 0x9c, 0x91, 0x84, 0xf5, 0x43, 0xfb, 0xdb, 0x4e, 0x33,
	0x2e, 0x39, 0xaa, 0xc5, 0x39, 0xf3, 0x36, 0x09, 0xef, 0xe5, 0xa7, 0x94, 0xc9, 0x50, 0xd0, 0xec,
	0x2c, 0xe9, 0x51, 0xa3, 0xf4, 0xee, 0xf7, 0x39, 0xef, 0x0f, 0x68, 0x88, 0xd3, 0x24, 0xc4, 0x8c,
	0x71, 0x89, 0x65, 0xc2, 0x99, 0xb0, 0xda, 0x96, 0xfe, 0xe9, 0xed, 0xf4, 0x29, 0xdb, 0x11, 0xaf,
	0x70, 0xbf, 0x4f, 0xb3, 0x90, 0xa7, 0xda, 0x62, 0xda, 0x3a, 0x78, 0x01, 0xeb, 0x5f, 0x99, 0xc8,
	0x5f, 0x66, 0x14, 0x4b, 0x7a, 0x84, 0x87, 0x03, 0x8e, 0x09, 0x7a, 0x00, 0x90, 0x10, 0xca, 0x64,
	0x12, 0x27, 0x34, 0x73, 0x1d, 0xdf, 0x79, 0xb8, 0xd2, 0xa9, 0x48, 0xd0, 0x87, 0xb0, 0x40, 0xb0,
	0xc4, 0xee, 0x2d, 0xdf, 0x79, 0xd8, 0xd8, 0x5b, 0x6b, 0xc7, 0x39, 0x6b, 0x5b, 0x47, 0xcf, 0xb1,
	0xc4, 0x1d, 0xad, 0x0d, 0x7e, 0x2e, 0xbd, 0x7f, 0x9f, 0x92, 0x6b, 0x78, 0xff, 0x00, 0x56, 0x71,
	0x3f, 0xa3, 0x54, 0x1d, 0xbe, 0x9b, 0x10, 0x1d, 0x65, 0xa5, 0xd3, 0x28, 0x65, 0x07, 0xa4, 0x04,
	0xa8, 0xbd, 0x15, 0xe0, 0x25, 0xdc, 0xb1, 0xc2, 0x0e, 0x15, 0x29, 0x67, 0x82, 0xa2, 0x5d, 0x58,
	0x3a, 0xa1, 0x98, 0xd8, 0xb8, 0x8d, 0x3d, 0xb7, 0x5d, 0xa4, 0xb9, 0x5d, 0xd8, 0x7c, 0xad, 0xf5,
	0x1d, 0x6b, 0x87, 0x5a, 0x13, 0x67, 0x75, 0xab, 0xa1, 0x8a, 0x1d, 0x95, 0x90, 0x39, 0xdc, 0xb3,
	0xca, 0xc3, 0x44, 0xc8, 0x1b, 0x09, 0x5b, 0x9b, 0x23, 0xec, 0x21, 0xd4, 0x3b, 0xf4, 0x65, 0x4e,
	0x85, 0xbc, 0x81, 0xec, 0x06, 0x29, 0xdc, 0xdd, 0xa7, 0xf2, 0x07, 0x9a, 0x89, 0x84, 0xb3, 0x79,
	0xfd, 0xba, 0x50, 0x3f, 0x33, 0x3b, 0xac, 0xcb, 0x62, 0x39, 0x15, 0xb1, 0x36, 0x1d, 0x71, 0x17,
	0x9a, 0xfb, 0x54, 0x9a, 0x94, 0xcd, 0x15, 0x2e, 0xf8, 0x16, 0x36, 0xec, 0x8e, 0x9b, 0xe2, 0x0c,
	0xfe, 0xae, 0x41, 0xa3, 0x52, 0x44, 0x53, 0xdc, 0xce, 0x74, 0x1d, 0x6e, 0xc2, 0x12, 0x3e, 0xe5,
	0x39, 0x93, 0xd6, 0x97, 0x5d, 0xa1, 0x35, 0xa8, 0xe1, 0x34, 0xb3, 0x27, 0x55, 0x7f, 0x11, 0x52,
	0xf7, 0x39, 0x14, 0xee, 0x82, 0x16, 0xe9, 0xff, 0xca, 0x2a, 0xa6, 0xd4, 0x5d, 0x34, 0x56, 0x31,
	0xa5, 0xa8, 0x05, 0x28, 0xa3, 0x29, 0x1e, 0xea, 0x90, 0x24, 0xa7, 0x5d, 0xd5, 0x37, 0x6e, 0x5d,
	0x1b, 0xac, 0x95, 0x9a, 0xe7, 0xb9, 0xba, 0x7b, 0x8a, 0x3e, 0x83, 0x77, 0xc6, 0xd6, 0xbc, 0xd7,
	0xcb, 0xb3, 0x8c, 0x12, 0xb3, 0x65, 0x59, 0x6f, 0xd9, 0x28, 0xd5, 0xdf, 0x58, 0xad, 0xde, 0xf7,
	0x08, 0xc6, 0xbe, 0xba, 0x96, 0x7f, 0x45, 0x6f, 0xb8, 0x53, 0xca, 0x23, 0x73, 0x10, 0x0f, 0x96,
	0xf5, 0x46, 0xd6, 0x1b, 0xba, 0xa0, 0x4d, 0xca, 0x35, 0x7a, 0x1f, 0x1a, 0x2c, 0x96, 0x5d, 0x4c,
	0x48, 0x46, 0x85, 0x70, 0x1b, 0x26, 0xd5, 0x2c, 0x96, 0x91, 0x91, 0xa8, 0xd3, 0x94, 0x67, 0xa1,
	0x12, 0x27, 0x03, 0xa1, 0xd2, 0xb8, 0x6a, 0x4e, 0x53, 0x9c, 0xc5, 0x28, 0x0e, 0x08, 0x7a, 0x0f,
	0x56, 0xd4, 0x33, 0x48, 0x33, 0x65, 0x74, 0xdb, 0xc4, 0x32, 0x82, 0x03, 0xa2, 0x62, 0x1d, 0xf3,
	0x2c, 0xe3, 0xaf, 0x8c, 0xba, 0x69, 0x62, 0x15, 0xa2, 0x03, 0x12, 0xfc, 0x54, 0x36, 0x5e, 0xb5,
	0x3d, 0xd0, 0x36, 0xd4, 0xed, 0xdb, 0x6a, 0x3b, 0x6f, 0xfa, 0xad, 0x28, 0x0c, 0xd0, 0xa7, 0x00,
	0x22, 0xe9, 0x33, 0x2c, 0xf3, 0x8c, 0x0a, 0xdb, 0x78, 0x1b, 0x55, 0xf3, 0xef, 0x0a, 0x6d, 0xa7,
	0x62, 0x18, 0xfc, 0xe1, 0xc0, 0xda, 0x45, 0x03, 0xb4, 0x0e, 0x8b, 0x67, 0x78, 0x50, 0x16, 0x8d,
	0x59, 0xa0, 0x1d, 0x40, 0x3c, 0x97, 0xea, 0x82, 0x48, 0xb7, 0xf4, 0x60, 0x4b, 0xe7, 0x6e, 0xa1,
	0x19, 0x3b, 0xf1, 0x60, 0xd9, 0x14, 0xae, 0x1c, 0xda, 0x52, 0x2a, 0xd7, 0x68, 0x0b, 0x9a, 0xca,
	0x03, 0x25, 0xdd, 0xa2, 0x9a, 0x4d, 0x65, 0xdd, 0x36, 0x52, 0xdb, 0x14, 0x7b, 0xff, 0x2d, 0x43,
	0xb3, 0x80, 0x33, 0x63, 0x04, 0xfd, 0xea, 0xc0, 0x92, 0x79, 0xee, 0xd1, 0xbb, 0xd5, 0xd3, 0x4d,
	0x8c, 0x00, 0x6f, 0xfd, 0xb2, 0x17, 0x27, 0x38, 0x1c, 0x45, 0x0f, 0xbc, 0xfb, 0x11, 0x21, 0xc2,
	0xc7, 0xbe, 0xcd, 0x9b, 0x2f, 0xb9, 0x8f, 0xfd, 0xe2, 0x3d, 0xfb, 0xe5, 0x9f, 0x7f, 0xdf, 0xdc,
	0xda, 0x0a, 0xfc, 0xb0, 0x10, 0x88, 0xf0, 0x7c, 0xdc, 0x7d, 0xaf, 0x8b, 0x49, 0x27, 0x3e, 0x77,
	0xb6, 0xd1, 0x9f, 0x0e, 0x2c, 0x99, 0xd1, 0x30, 0x49, 0x32, 0x31, 0x2e, 0x66, 0x90, 0x9c, 0x8c,
	0xa2, 0x8f, 0xbd, 0x47, 0xc6, 0xb2, 0x0a, 0x53, 0x76, 0xa7, 0x9f, 0xb0, 0x8b, 0x58, 0x4f, 0xbc,
	0xf0, 0x2a, 0xac, 0xf0, 0xbc, 0xda, 0xf2, 0xaf, 0x15, 0xe5, 0x6f, 0x0e, 0x2c, 0xa8, 0x3b, 0x41,
	0xab, 0x1a, 0xc4, 0xbe, 0x33, 0x33, 0xb0, 0x92, 0x51, 0xf4, 0x91, 0xb7, 0xa5, 0xcc, 0xc5, 0x5c,
	0x48, 0x4f, 0x83, 0x27, 0xd7, 0x44, 0x0a, 0xd5, 0x05, 0x2b, 0xae, 0x37, 0x0e, 0xd4, 0xf6, 0xa9,
	0x9c, 0x0b, 0xeb, 0x64, 0x14, 0xb5, 0xbd, 0xd6, 0x3e, 0x95, 0x97, 0x66, 0x8a, 0xc7, 0x3e, 0xf6,
	0x07, 0x2a, 0x91, 0x72, 0x92, 0xee, 0x31, 0xba, 0x6e, 0xc2, 0xd0, 0x5f, 0x0e, 0xc0, 0x78, 0x78,
	0xa0, 0x4d, 0x8d, 0x33, 0x35, 0x4d, 0x66, 0x60, 0x8a, 0x51, 0x14, 0x7a, 0x3b, 0x6f, 0xc5, 0x2c,
	0x70, 0x7c, 0x5b, 0xf6, 0x9a, 0xf3, 0x0b, 0xf4, 0x74, 0x16, 0xe7, 0xb9, 0xb5, 0x9b, 0x4d, 0xfc,
	0xbb, 0x03, 0x75, 0x3b, 0x4a, 0xd0, 0xbd, 0x02, 0xb7, 0x32, 0x8a, 0xbc, 0x89, 0xe1, 0x5b, 0x1d,
	0xeb, 0xc1, 0x8b, 0x51, 0xf4, 0xd8, 0x0b, 0x35, 0xef, 0x60, 0x30, 0x4d, 0x2c, 0x66, 0x67, 0x36,
	0x40, 0x57, 0x76, 0x88, 0x4a, 0x65, 0x73, 0x72, 0xc6, 0x21, 0xaf, 0xca, 0x77, 0x21, 0xa5, 0xb3,
	0x31, 0xfb, 0xa3, 0x68, 0xcf, 0xdb, 0xbd, 0x0a, 0xf3, 0xd2, 0xcc, 0xb6, 0xd0, 0xf6, 0xfc, 0x99,
	0x7d, 0xe6, 0x43, 0xbd, 0xc7, 0x4f, 0x15, 0xc7, 0xb3, 0x55, 0x0b, 0x72, 0xa4, 0x3e, 0x34, 0x8f,
	0x9c, 0x1f, 0x17, 0xe3, 0x9c, 0xa5, 0xc7, 0xc7, 0x4b, 0xfa, 0xc3, 0xf3, 0x93, 0xff, 0x03, 0x00,
	0x00, 0xff, 0xff, 0xe6, 0x34, 0x7c, 0x81, 0xfa, 0x0a, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FundingServiceClient is the client API for FundingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FundingServiceClient interface {
	Create(ctx context.Context, in *FundingCreatePayload, opts ...grpc.CallOption) (*FundingResponse, error)
	Update(ctx context.Context, in *FundingUpdatePayload, opts ...grpc.CallOption) (*FundingResponse, error)
	Sign(ctx context.Context, in *Request, opts ...grpc.CallOption) (*FundingResponse, error)
	Get(ctx context.Context, in *Request, opts ...grpc.CallOption) (*FundingResponse, error)
	GetVersion(ctx context.Context, in *GetVersionRequest, opts ...grpc.CallOption) (*FundingResponse, error)
	GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*FundingListResponse, error)
	GetListVersion(ctx context.Context, in *GetListVersionRequest, opts ...grpc.CallOption) (*FundingListResponse, error)
}

type fundingServiceClient struct {
	cc *grpc.ClientConn
}

func NewFundingServiceClient(cc *grpc.ClientConn) FundingServiceClient {
	return &fundingServiceClient{cc}
}

func (c *fundingServiceClient) Create(ctx context.Context, in *FundingCreatePayload, opts ...grpc.CallOption) (*FundingResponse, error) {
	out := new(FundingResponse)
	err := c.cc.Invoke(ctx, "/fun.FundingService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fundingServiceClient) Update(ctx context.Context, in *FundingUpdatePayload, opts ...grpc.CallOption) (*FundingResponse, error) {
	out := new(FundingResponse)
	err := c.cc.Invoke(ctx, "/fun.FundingService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fundingServiceClient) Sign(ctx context.Context, in *Request, opts ...grpc.CallOption) (*FundingResponse, error) {
	out := new(FundingResponse)
	err := c.cc.Invoke(ctx, "/fun.FundingService/Sign", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fundingServiceClient) Get(ctx context.Context, in *Request, opts ...grpc.CallOption) (*FundingResponse, error) {
	out := new(FundingResponse)
	err := c.cc.Invoke(ctx, "/fun.FundingService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fundingServiceClient) GetVersion(ctx context.Context, in *GetVersionRequest, opts ...grpc.CallOption) (*FundingResponse, error) {
	out := new(FundingResponse)
	err := c.cc.Invoke(ctx, "/fun.FundingService/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fundingServiceClient) GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*FundingListResponse, error) {
	out := new(FundingListResponse)
	err := c.cc.Invoke(ctx, "/fun.FundingService/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fundingServiceClient) GetListVersion(ctx context.Context, in *GetListVersionRequest, opts ...grpc.CallOption) (*FundingListResponse, error) {
	out := new(FundingListResponse)
	err := c.cc.Invoke(ctx, "/fun.FundingService/GetListVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FundingServiceServer is the server API for FundingService service.
type FundingServiceServer interface {
	Create(context.Context, *FundingCreatePayload) (*FundingResponse, error)
	Update(context.Context, *FundingUpdatePayload) (*FundingResponse, error)
	Sign(context.Context, *Request) (*FundingResponse, error)
	Get(context.Context, *Request) (*FundingResponse, error)
	GetVersion(context.Context, *GetVersionRequest) (*FundingResponse, error)
	GetList(context.Context, *GetListRequest) (*FundingListResponse, error)
	GetListVersion(context.Context, *GetListVersionRequest) (*FundingListResponse, error)
}

func RegisterFundingServiceServer(s *grpc.Server, srv FundingServiceServer) {
	s.RegisterService(&_FundingService_serviceDesc, srv)
}

func _FundingService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FundingCreatePayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FundingServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fun.FundingService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FundingServiceServer).Create(ctx, req.(*FundingCreatePayload))
	}
	return interceptor(ctx, in, info, handler)
}

func _FundingService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FundingUpdatePayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FundingServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fun.FundingService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FundingServiceServer).Update(ctx, req.(*FundingUpdatePayload))
	}
	return interceptor(ctx, in, info, handler)
}

func _FundingService_Sign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FundingServiceServer).Sign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fun.FundingService/Sign",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FundingServiceServer).Sign(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _FundingService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FundingServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fun.FundingService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FundingServiceServer).Get(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _FundingService_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FundingServiceServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fun.FundingService/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FundingServiceServer).GetVersion(ctx, req.(*GetVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FundingService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FundingServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fun.FundingService/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FundingServiceServer).GetList(ctx, req.(*GetListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FundingService_GetListVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FundingServiceServer).GetListVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fun.FundingService/GetListVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FundingServiceServer).GetListVersion(ctx, req.(*GetListVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FundingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fun.FundingService",
	HandlerType: (*FundingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _FundingService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _FundingService_Update_Handler,
		},
		{
			MethodName: "Sign",
			Handler:    _FundingService_Sign_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _FundingService_Get_Handler,
		},
		{
			MethodName: "GetVersion",
			Handler:    _FundingService_GetVersion_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _FundingService_GetList_Handler,
		},
		{
			MethodName: "GetListVersion",
			Handler:    _FundingService_GetListVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "funding/funding.proto",
}