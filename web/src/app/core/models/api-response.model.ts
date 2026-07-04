export interface ApiFieldError {
  field: string;
  message: string;
}

export interface ApiResponse<T> {
  path: string;
  timestamp: string;
  status: string;
  code: string;
  message: string;
  result: T;
  errors: ApiFieldError[] | null;
}

export interface AuthUser {
  userID: number;
  username: string;
  email: string;
  fullName: string;
  jobTitle: string;
  photoUrl: string;
  isAdmin: number;
  flagActive: number;
}

export interface LoginResult {
  accessToken: string;
  tokenType: string;
  expiresIn: number;
  user: AuthUser;
}
