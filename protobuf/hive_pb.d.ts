// package: hive
// file: github.com/benka-me/hive/protobuf/hive.proto

import * as jspb from "google-protobuf";
import * as github_com_gogo_protobuf_gogoproto_gogo_pb from "../../../../github.com/gogo/protobuf/gogoproto/gogo_pb";
import * as github_com_benka_me_cell_user_protobuf_user_pb from "../../../../github.com/benka-me/cell-user/protobuf/user_pb";

export class Bee extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  getPkgname(): string;
  setPkgname(value: string): void;

  getPkgnamecamel(): string;
  setPkgnamecamel(value: string): void;

  getRepo(): string;
  setRepo(value: string): void;

  getAuthor(): string;
  setAuthor(value: string): void;

  getAuthoremail(): string;
  setAuthoremail(value: string): void;

  getPort(): number;
  setPort(value: number): void;

  getPublic(): boolean;
  setPublic(value: boolean): void;

  getLicense(): string;
  setLicense(value: string): void;

  getDescription(): string;
  setDescription(value: string): void;

  getKeywords(): string;
  setKeywords(value: string): void;

  getTag(): string;
  setTag(value: string): void;

  getDevlang(): DevLangMap[keyof DevLangMap];
  setDevlang(value: DevLangMap[keyof DevLangMap]): void;

  hasLanguages(): boolean;
  clearLanguages(): void;
  getLanguages(): Languages | undefined;
  setLanguages(value?: Languages): void;

  getIsgateway(): boolean;
  setIsgateway(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Bee.AsObject;
  static toObject(includeInstance: boolean, msg: Bee): Bee.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Bee, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Bee;
  static deserializeBinaryFromReader(message: Bee, reader: jspb.BinaryReader): Bee;
}

export namespace Bee {
  export type AsObject = {
    name: string,
    pkgname: string,
    pkgnamecamel: string,
    repo: string,
    author: string,
    authoremail: string,
    port: number,
    pb_public: boolean,
    license: string,
    description: string,
    keywords: string,
    tag: string,
    devlang: DevLangMap[keyof DevLangMap],
    languages?: Languages.AsObject,
    isgateway: boolean,
  }
}

export class Bees extends jspb.Message {
  clearBeesList(): void;
  getBeesList(): Array<Bee>;
  setBeesList(value: Array<Bee>): void;
  addBees(value?: Bee, index?: number): Bee;

  getStatusmessage(): string;
  setStatusmessage(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Bees.AsObject;
  static toObject(includeInstance: boolean, msg: Bees): Bees.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Bees, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Bees;
  static deserializeBinaryFromReader(message: Bees, reader: jspb.BinaryReader): Bees;
}

export namespace Bees {
  export type AsObject = {
    beesList: Array<Bee.AsObject>,
    statusmessage: string,
  }
}

export class BeeReq extends jspb.Message {
  hasToken(): boolean;
  clearToken(): void;
  getToken(): github_com_benka_me_cell_user_protobuf_user_pb.Token | undefined;
  setToken(value?: github_com_benka_me_cell_user_protobuf_user_pb.Token): void;

  getBeename(): string;
  setBeename(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): BeeReq.AsObject;
  static toObject(includeInstance: boolean, msg: BeeReq): BeeReq.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: BeeReq, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): BeeReq;
  static deserializeBinaryFromReader(message: BeeReq, reader: jspb.BinaryReader): BeeReq;
}

export namespace BeeReq {
  export type AsObject = {
    token?: github_com_benka_me_cell_user_protobuf_user_pb.Token.AsObject,
    beename: string,
  }
}

export class BeesReq extends jspb.Message {
  hasToken(): boolean;
  clearToken(): void;
  getToken(): github_com_benka_me_cell_user_protobuf_user_pb.Token | undefined;
  setToken(value?: github_com_benka_me_cell_user_protobuf_user_pb.Token): void;

  clearBeenamesList(): void;
  getBeenamesList(): Array<string>;
  setBeenamesList(value: Array<string>): void;
  addBeenames(value: string, index?: number): string;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): BeesReq.AsObject;
  static toObject(includeInstance: boolean, msg: BeesReq): BeesReq.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: BeesReq, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): BeesReq;
  static deserializeBinaryFromReader(message: BeesReq, reader: jspb.BinaryReader): BeesReq;
}

export namespace BeesReq {
  export type AsObject = {
    token?: github_com_benka_me_cell_user_protobuf_user_pb.Token.AsObject,
    beenamesList: Array<string>,
  }
}

export class Dep extends jspb.Message {
  getPort(): number;
  setPort(value: number): void;

  getDev(): string;
  setDev(value: string): void;

  getProd(): string;
  setProd(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Dep.AsObject;
  static toObject(includeInstance: boolean, msg: Dep): Dep.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Dep, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Dep;
  static deserializeBinaryFromReader(message: Dep, reader: jspb.BinaryReader): Dep;
}

export namespace Dep {
  export type AsObject = {
    port: number,
    dev: string,
    prod: string,
  }
}

export class Hive extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  getPkgname(): string;
  setPkgname(value: string): void;

  getPkgnamecamel(): string;
  setPkgnamecamel(value: string): void;

  getRepo(): string;
  setRepo(value: string): void;

  getAuthor(): string;
  setAuthor(value: string): void;

  getPublic(): boolean;
  setPublic(value: boolean): void;

  getDepsMap(): jspb.Map<string, Dep>;
  clearDepsMap(): void;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Hive.AsObject;
  static toObject(includeInstance: boolean, msg: Hive): Hive.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Hive, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Hive;
  static deserializeBinaryFromReader(message: Hive, reader: jspb.BinaryReader): Hive;
}

export namespace Hive {
  export type AsObject = {
    name: string,
    pkgname: string,
    pkgnamecamel: string,
    repo: string,
    author: string,
    pb_public: boolean,
    depsMap: Array<[string, Dep.AsObject]>,
  }
}

export class Languages extends jspb.Message {
  hasGo(): boolean;
  clearGo(): void;
  getGo(): Go | undefined;
  setGo(value?: Go): void;

  hasJavascript(): boolean;
  clearJavascript(): void;
  getJavascript(): Javascript | undefined;
  setJavascript(value?: Javascript): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Languages.AsObject;
  static toObject(includeInstance: boolean, msg: Languages): Languages.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Languages, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Languages;
  static deserializeBinaryFromReader(message: Languages, reader: jspb.BinaryReader): Languages;
}

export namespace Languages {
  export type AsObject = {
    go?: Go.AsObject,
    javascript?: Javascript.AsObject,
  }
}

export class LanguageSetup extends jspb.Message {
  getActive(): boolean;
  setActive(value: boolean): void;

  getRepo(): string;
  setRepo(value: string): void;

  clearFilesList(): void;
  getFilesList(): Array<string>;
  setFilesList(value: Array<string>): void;
  addFiles(value: string, index?: number): string;

  getPkgname(): string;
  setPkgname(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LanguageSetup.AsObject;
  static toObject(includeInstance: boolean, msg: LanguageSetup): LanguageSetup.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: LanguageSetup, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LanguageSetup;
  static deserializeBinaryFromReader(message: LanguageSetup, reader: jspb.BinaryReader): LanguageSetup;
}

export namespace LanguageSetup {
  export type AsObject = {
    active: boolean,
    repo: string,
    filesList: Array<string>,
    pkgname: string,
  }
}

export class Go extends jspb.Message {
  hasSetup(): boolean;
  clearSetup(): void;
  getSetup(): LanguageSetup | undefined;
  setSetup(value?: LanguageSetup): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Go.AsObject;
  static toObject(includeInstance: boolean, msg: Go): Go.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Go, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Go;
  static deserializeBinaryFromReader(message: Go, reader: jspb.BinaryReader): Go;
}

export namespace Go {
  export type AsObject = {
    setup?: LanguageSetup.AsObject,
  }
}

export class Javascript extends jspb.Message {
  hasSetup(): boolean;
  clearSetup(): void;
  getSetup(): LanguageSetup | undefined;
  setSetup(value?: LanguageSetup): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Javascript.AsObject;
  static toObject(includeInstance: boolean, msg: Javascript): Javascript.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Javascript, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Javascript;
  static deserializeBinaryFromReader(message: Javascript, reader: jspb.BinaryReader): Javascript;
}

export namespace Javascript {
  export type AsObject = {
    setup?: LanguageSetup.AsObject,
  }
}

export class Python extends jspb.Message {
  hasSetup(): boolean;
  clearSetup(): void;
  getSetup(): LanguageSetup | undefined;
  setSetup(value?: LanguageSetup): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Python.AsObject;
  static toObject(includeInstance: boolean, msg: Python): Python.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Python, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Python;
  static deserializeBinaryFromReader(message: Python, reader: jspb.BinaryReader): Python;
}

export namespace Python {
  export type AsObject = {
    setup?: LanguageSetup.AsObject,
  }
}

export class Java extends jspb.Message {
  hasSetup(): boolean;
  clearSetup(): void;
  getSetup(): LanguageSetup | undefined;
  setSetup(value?: LanguageSetup): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Java.AsObject;
  static toObject(includeInstance: boolean, msg: Java): Java.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Java, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Java;
  static deserializeBinaryFromReader(message: Java, reader: jspb.BinaryReader): Java;
}

export namespace Java {
  export type AsObject = {
    setup?: LanguageSetup.AsObject,
  }
}

export class CPP extends jspb.Message {
  hasSetup(): boolean;
  clearSetup(): void;
  getSetup(): LanguageSetup | undefined;
  setSetup(value?: LanguageSetup): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CPP.AsObject;
  static toObject(includeInstance: boolean, msg: CPP): CPP.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CPP, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CPP;
  static deserializeBinaryFromReader(message: CPP, reader: jspb.BinaryReader): CPP;
}

export namespace CPP {
  export type AsObject = {
    setup?: LanguageSetup.AsObject,
  }
}

export class CSharp extends jspb.Message {
  hasSetup(): boolean;
  clearSetup(): void;
  getSetup(): LanguageSetup | undefined;
  setSetup(value?: LanguageSetup): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CSharp.AsObject;
  static toObject(includeInstance: boolean, msg: CSharp): CSharp.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CSharp, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CSharp;
  static deserializeBinaryFromReader(message: CSharp, reader: jspb.BinaryReader): CSharp;
}

export namespace CSharp {
  export type AsObject = {
    setup?: LanguageSetup.AsObject,
  }
}

export class ObjectiveC extends jspb.Message {
  hasSetup(): boolean;
  clearSetup(): void;
  getSetup(): LanguageSetup | undefined;
  setSetup(value?: LanguageSetup): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ObjectiveC.AsObject;
  static toObject(includeInstance: boolean, msg: ObjectiveC): ObjectiveC.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ObjectiveC, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ObjectiveC;
  static deserializeBinaryFromReader(message: ObjectiveC, reader: jspb.BinaryReader): ObjectiveC;
}

export namespace ObjectiveC {
  export type AsObject = {
    setup?: LanguageSetup.AsObject,
  }
}

export class Ruby extends jspb.Message {
  hasSetup(): boolean;
  clearSetup(): void;
  getSetup(): LanguageSetup | undefined;
  setSetup(value?: LanguageSetup): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Ruby.AsObject;
  static toObject(includeInstance: boolean, msg: Ruby): Ruby.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Ruby, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Ruby;
  static deserializeBinaryFromReader(message: Ruby, reader: jspb.BinaryReader): Ruby;
}

export namespace Ruby {
  export type AsObject = {
    setup?: LanguageSetup.AsObject,
  }
}

export class Dart extends jspb.Message {
  hasSetup(): boolean;
  clearSetup(): void;
  getSetup(): LanguageSetup | undefined;
  setSetup(value?: LanguageSetup): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Dart.AsObject;
  static toObject(includeInstance: boolean, msg: Dart): Dart.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Dart, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Dart;
  static deserializeBinaryFromReader(message: Dart, reader: jspb.BinaryReader): Dart;
}

export namespace Dart {
  export type AsObject = {
    setup?: LanguageSetup.AsObject,
  }
}

export class PHP extends jspb.Message {
  hasSetup(): boolean;
  clearSetup(): void;
  getSetup(): LanguageSetup | undefined;
  setSetup(value?: LanguageSetup): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PHP.AsObject;
  static toObject(includeInstance: boolean, msg: PHP): PHP.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PHP, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PHP;
  static deserializeBinaryFromReader(message: PHP, reader: jspb.BinaryReader): PHP;
}

export namespace PHP {
  export type AsObject = {
    setup?: LanguageSetup.AsObject,
  }
}

export class PushBee extends jspb.Message {
  hasBee(): boolean;
  clearBee(): void;
  getBee(): Bee | undefined;
  setBee(value?: Bee): void;

  hasToken(): boolean;
  clearToken(): void;
  getToken(): github_com_benka_me_cell_user_protobuf_user_pb.Token | undefined;
  setToken(value?: github_com_benka_me_cell_user_protobuf_user_pb.Token): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PushBee.AsObject;
  static toObject(includeInstance: boolean, msg: PushBee): PushBee.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PushBee, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PushBee;
  static deserializeBinaryFromReader(message: PushBee, reader: jspb.BinaryReader): PushBee;
}

export namespace PushBee {
  export type AsObject = {
    bee?: Bee.AsObject,
    token?: github_com_benka_me_cell_user_protobuf_user_pb.Token.AsObject,
  }
}

export class PushHive extends jspb.Message {
  hasHive(): boolean;
  clearHive(): void;
  getHive(): Hive | undefined;
  setHive(value?: Hive): void;

  hasToken(): boolean;
  clearToken(): void;
  getToken(): github_com_benka_me_cell_user_protobuf_user_pb.Token | undefined;
  setToken(value?: github_com_benka_me_cell_user_protobuf_user_pb.Token): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PushHive.AsObject;
  static toObject(includeInstance: boolean, msg: PushHive): PushHive.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PushHive, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PushHive;
  static deserializeBinaryFromReader(message: PushHive, reader: jspb.BinaryReader): PushHive;
}

export namespace PushHive {
  export type AsObject = {
    hive?: Hive.AsObject,
    token?: github_com_benka_me_cell_user_protobuf_user_pb.Token.AsObject,
  }
}

export class PushBeeRes extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PushBeeRes.AsObject;
  static toObject(includeInstance: boolean, msg: PushBeeRes): PushBeeRes.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PushBeeRes, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PushBeeRes;
  static deserializeBinaryFromReader(message: PushBeeRes, reader: jspb.BinaryReader): PushBeeRes;
}

export namespace PushBeeRes {
  export type AsObject = {
    id: string,
  }
}

export class PushHiveRes extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PushHiveRes.AsObject;
  static toObject(includeInstance: boolean, msg: PushHiveRes): PushHiveRes.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PushHiveRes, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PushHiveRes;
  static deserializeBinaryFromReader(message: PushHiveRes, reader: jspb.BinaryReader): PushHiveRes;
}

export namespace PushHiveRes {
  export type AsObject = {
    id: string,
  }
}

export interface DevLangMap {
  GO_: 0;
  JAVASCRIPT_: 1;
  PYTHON_: 2;
  JAVA_: 3;
  CPP_: 4;
  CSHARP_: 5;
  OBJECTIVEC_: 6;
  RUBY_: 7;
  DART_: 8;
  PHP_: 9;
}

export const DevLang: DevLangMap;

