export interface PeridotUser {
  sub: string;
  email: string;
}

declare global {
  interface Window {
    __peridot_prefix__: string;
    __peridot_user__: PeridotUser;
  }
}
