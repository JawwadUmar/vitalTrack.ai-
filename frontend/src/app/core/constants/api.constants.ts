export const API_CONSTANTS = {
  BASE_URL: 'http://localhost:8081/api/v1',
  get LOGIN_URL() {
    return `${this.BASE_URL}/users/login`;
  },
  get SIGNUP_URL() {
    return `${this.BASE_URL}/users/signup`;
  },
  get DOCUMENTS_CALENDAR_URL() {
    return `${this.BASE_URL}/documents/calendar`;
  },
  get FILES_UPLOAD_URL() {
    return `${this.BASE_URL}/files/upload`;
  },
  get DOCUMENTS_URL() {
    return `${this.BASE_URL}/documents`;
  }
};
