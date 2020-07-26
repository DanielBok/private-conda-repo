import * as T from "./types";

class ChannelLocalStorage {
  private itemKey = "PCR_CHANNEL_INFO";

  public load() {
    const item = localStorage.getItem(this.itemKey);
    if (item !== null) {
      return JSON.parse(item) as T.Channel;
    }
    return null;
  }

  public save(channel: T.Channel) {
    localStorage.setItem(this.itemKey, JSON.stringify(channel));
  }

  public clear() {
    localStorage.removeItem(this.itemKey);
  }
}

export const ChannelStorage = new ChannelLocalStorage();
