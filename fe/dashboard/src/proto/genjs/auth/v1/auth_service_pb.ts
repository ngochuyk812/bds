// @generated by protoc-gen-es v2.3.0 with parameter "target=ts"
// @generated from file auth/v1/auth_service.proto (package auth.v1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenService } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, serviceDesc } from "@bufbuild/protobuf/codegenv1";
import type { CreateSiteRequestSchema, CreateSiteResponseSchema, DeleteSiteRequestSchema, DeleteSiteResponseSchema, FetchSitesRequestSchema, FetchSitesResponseSchema, UpdateSiteRequestSchema, UpdateSiteResponseSchema } from "./site_pb";
import { file_auth_v1_site } from "./site_pb";
import type { LoginRequestSchema, LoginResponseSchema, LogoutRequestSchema, LogoutResponseSchema, RefreshTokenRequestSchema, RefreshTokenResponseSchema, SignUpRequestSchema, SignUpResponseSchema, VerifySignUpRequestSchema, VerifySignUpResponseSchema } from "./auth_pb";
import { file_auth_v1_auth } from "./auth_pb";
import type { GetProfileRequestSchema, GetProfileResponseSchema, UpdateProfileRequestSchema, UpdateProfileResponseSchema } from "./user_pb";
import { file_auth_v1_user } from "./user_pb";

/**
 * Describes the file auth/v1/auth_service.proto.
 */
export const file_auth_v1_auth_service: GenFile = /*@__PURE__*/
  fileDesc("ChphdXRoL3YxL2F1dGhfc2VydmljZS5wcm90bxIHYXV0aC52MTKKBgoLQXV0aFNlcnZpY2USRQoKRmV0Y2hTaXRlcxIaLmF1dGgudjEuRmV0Y2hTaXRlc1JlcXVlc3QaGy5hdXRoLnYxLkZldGNoU2l0ZXNSZXNwb25zZRJFCgpDcmVhdGVTaXRlEhouYXV0aC52MS5DcmVhdGVTaXRlUmVxdWVzdBobLmF1dGgudjEuQ3JlYXRlU2l0ZVJlc3BvbnNlEkUKClVwZGF0ZVNpdGUSGi5hdXRoLnYxLlVwZGF0ZVNpdGVSZXF1ZXN0GhsuYXV0aC52MS5VcGRhdGVTaXRlUmVzcG9uc2USRQoKRGVsZXRlU2l0ZRIaLmF1dGgudjEuRGVsZXRlU2l0ZVJlcXVlc3QaGy5hdXRoLnYxLkRlbGV0ZVNpdGVSZXNwb25zZRI2CgVMb2dpbhIVLmF1dGgudjEuTG9naW5SZXF1ZXN0GhYuYXV0aC52MS5Mb2dpblJlc3BvbnNlEjkKBlNpZ25VcBIWLmF1dGgudjEuU2lnblVwUmVxdWVzdBoXLmF1dGgudjEuU2lnblVwUmVzcG9uc2USSwoMVmVyaWZ5U2lnblVwEhwuYXV0aC52MS5WZXJpZnlTaWduVXBSZXF1ZXN0Gh0uYXV0aC52MS5WZXJpZnlTaWduVXBSZXNwb25zZRJLCgxSZWZyZXNoVG9rZW4SHC5hdXRoLnYxLlJlZnJlc2hUb2tlblJlcXVlc3QaHS5hdXRoLnYxLlJlZnJlc2hUb2tlblJlc3BvbnNlEjkKBkxvZ291dBIWLmF1dGgudjEuTG9nb3V0UmVxdWVzdBoXLmF1dGgudjEuTG9nb3V0UmVzcG9uc2USRQoKR2V0UHJvZmlsZRIaLmF1dGgudjEuR2V0UHJvZmlsZVJlcXVlc3QaGy5hdXRoLnYxLkdldFByb2ZpbGVSZXNwb25zZRJQCg1VcGRhdGVQcm9maWxlEh0uYXV0aC52MS5VcGRhdGVQcm9maWxlUmVxdWVzdBoeLmF1dGgudjEuVXBkYXRlUHJvZmlsZVJlc3BvbnNlIgBCNVozZ2l0aHViLmNvbS9uZ29jaHV5azgxMi9wcm90by1iZHMvZ2VuL2F1dGgvdjE7YXV0aHYxYgZwcm90bzM", [file_auth_v1_site, file_auth_v1_auth, file_auth_v1_user]);

/**
 * @generated from service auth.v1.AuthService
 */
export const AuthService: GenService<{
  /**
   * @generated from rpc auth.v1.AuthService.FetchSites
   */
  fetchSites: {
    methodKind: "unary";
    input: typeof FetchSitesRequestSchema;
    output: typeof FetchSitesResponseSchema;
  },
  /**
   * @generated from rpc auth.v1.AuthService.CreateSite
   */
  createSite: {
    methodKind: "unary";
    input: typeof CreateSiteRequestSchema;
    output: typeof CreateSiteResponseSchema;
  },
  /**
   * @generated from rpc auth.v1.AuthService.UpdateSite
   */
  updateSite: {
    methodKind: "unary";
    input: typeof UpdateSiteRequestSchema;
    output: typeof UpdateSiteResponseSchema;
  },
  /**
   * @generated from rpc auth.v1.AuthService.DeleteSite
   */
  deleteSite: {
    methodKind: "unary";
    input: typeof DeleteSiteRequestSchema;
    output: typeof DeleteSiteResponseSchema;
  },
  /**
   * @generated from rpc auth.v1.AuthService.Login
   */
  login: {
    methodKind: "unary";
    input: typeof LoginRequestSchema;
    output: typeof LoginResponseSchema;
  },
  /**
   * @generated from rpc auth.v1.AuthService.SignUp
   */
  signUp: {
    methodKind: "unary";
    input: typeof SignUpRequestSchema;
    output: typeof SignUpResponseSchema;
  },
  /**
   * @generated from rpc auth.v1.AuthService.VerifySignUp
   */
  verifySignUp: {
    methodKind: "unary";
    input: typeof VerifySignUpRequestSchema;
    output: typeof VerifySignUpResponseSchema;
  },
  /**
   * @generated from rpc auth.v1.AuthService.RefreshToken
   */
  refreshToken: {
    methodKind: "unary";
    input: typeof RefreshTokenRequestSchema;
    output: typeof RefreshTokenResponseSchema;
  },
  /**
   * @generated from rpc auth.v1.AuthService.Logout
   */
  logout: {
    methodKind: "unary";
    input: typeof LogoutRequestSchema;
    output: typeof LogoutResponseSchema;
  },
  /**
   * @generated from rpc auth.v1.AuthService.GetProfile
   */
  getProfile: {
    methodKind: "unary";
    input: typeof GetProfileRequestSchema;
    output: typeof GetProfileResponseSchema;
  },
  /**
   * @generated from rpc auth.v1.AuthService.UpdateProfile
   */
  updateProfile: {
    methodKind: "unary";
    input: typeof UpdateProfileRequestSchema;
    output: typeof UpdateProfileResponseSchema;
  },
}> = /*@__PURE__*/
  serviceDesc(file_auth_v1_auth_service, 0);

