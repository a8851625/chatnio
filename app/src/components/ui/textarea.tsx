import * as React from "react";

import { cn } from "./lib/utils";
import { useMemo } from "react";

export interface TextareaProps
  extends React.TextareaHTMLAttributes<HTMLTextAreaElement> {}

const Textarea = React.forwardRef<HTMLTextAreaElement, TextareaProps>(
  ({ className, ...props }, ref) => {
    return (
      <textarea
        className={cn(
          "flex min-h-[80px] w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50",
          className,
        )}
        ref={ref}
        {...props}
      />
    );
  },
);
Textarea.displayName = "Textarea";

// FlexibleTextarea is a flexible rows textarea (current lines + 1)
export interface FlexibleTextareaProps extends TextareaProps {
  rows?: number;
  minRows?: number;
}

const FlexibleTextarea = React.forwardRef<
  HTMLTextAreaElement,
  FlexibleTextareaProps
>(({ rows = 1, ...props }, ref) => {
  const lines = useMemo(() => {
    const value = props.value?.toString() || "";
    const count = value.split("\n").length + 1;
    return Math.max(rows, count);
  }, [props.value]);

  return <Textarea ref={ref} rows={lines} {...props} />;
});

FlexibleTextarea.displayName = "FlexibleTextarea";

export { Textarea, FlexibleTextarea };
