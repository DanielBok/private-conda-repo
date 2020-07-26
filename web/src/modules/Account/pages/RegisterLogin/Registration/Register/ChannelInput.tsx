import { ChnAction, ChnApi } from "@/features/channel";
import { Form, Input } from "antd";
import React from "react";
import { useDispatch } from "react-redux";
import { useDebouncedCallback } from "use-debounce";
import Errors from "./Errors";
import { useDetails, useDisabled } from "./utils";

const ChannelInput = () => {
  const dispatch = useDispatch();
  const disabled = useDisabled();
  const [channel, errors, status] = useDetails("channel");
  const [checkAvailability] = useDebouncedCallback((name: string) => {
    if (name.length < 2) return;
    dispatch(ChnApi.isChannelAvailable(name));
  }, 500);

  return (
    <Form.Item
      validateStatus={status}
      hasFeedback
      help={<Errors errors={errors} />}
    >
      <Input
        autoFocus
        value={channel}
        placeholder="Username / Channel"
        onChange={({ target: { value: channel } }) => {
          channel = channel.trim().toLowerCase();
          dispatch(ChnAction.updateForm({ channel }));
          checkAvailability(channel);
        }}
        onKeyPress={(e) => {
          if (e.key === "Enter" && !disabled) dispatch(ChnApi.createChannel());
        }}
      />
    </Form.Item>
  );
};

export default ChannelInput;
