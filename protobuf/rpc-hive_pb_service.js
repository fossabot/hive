// package: RpcHive
// file: github.com/benka-me/hive/protobuf/rpc-hive.proto

var github_com_benka_me_hive_protobuf_rpc_hive_pb = require("../../../../github.com/benka-me/hive/protobuf/rpc-hive_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var Hive = (function () {
  function Hive() {}
  Hive.serviceName = "RpcHive.Hive";
  return Hive;
}());

exports.Hive = Hive;

function HiveClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

exports.HiveClient = HiveClient;

