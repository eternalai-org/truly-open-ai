export type TChatMessage = {
  id: string;
  is_reply: boolean;
  msg: string;
  name: string;
};

export interface IChatMessage {
  id: string;
  is_reply: boolean;
  msg: string;
  name: string;
}
