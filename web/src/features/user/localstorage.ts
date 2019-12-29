import * as UserType from "./types";

class UserLocalStorage {
  private itemKey = "PCR_USER_INFO";

  public load() {
    const item = localStorage.getItem(this.itemKey);
    if (item !== null) {
      return JSON.parse(item) as UserType.UserInfo;
    }
    return null;
  }

  public save(user: UserType.UserInfo) {
    localStorage.setItem(this.itemKey, JSON.stringify(user));
  }

  public clear() {
    localStorage.removeItem(this.itemKey);
  }
}

export const UserStorage = new UserLocalStorage();
