import { Button } from "./ui/button.tsx";
// import { useConversationActions, useMessages } from "@/store/chat.ts";
import { useConversationActions } from "@/store/chat.ts";
import { MessageSquarePlus } from "lucide-react";
// import Github from "@/components/ui/icons/Github.tsx";
// import { openWindow } from "@/utils/device.ts";

function ProjectLink() {
  // const messages = useMessages();

  const { toggle } = useConversationActions();

  return (
    <Button
      variant="outline"
      size="icon"
      onClick={async () => await toggle(-1)}
    >
      <MessageSquarePlus className={`h-4 w-4`} />
    </Button>
  );
}

export default ProjectLink;
