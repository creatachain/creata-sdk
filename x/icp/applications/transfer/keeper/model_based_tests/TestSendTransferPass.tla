------------------------- MODULE counterexample -------------------------

EXTENDS relay_tests

(* Initial state *)

State1 ==
TRUE
(* Transition 0 to State2 *)

State2 ==
/\ bank = <<
    [channel |-> "", id |-> "", port |-> ""], [denom |-> "",
      prefix0 |-> [channel |-> "", port |-> ""],
      prefix1 |-> [channel |-> "", port |-> ""]]
  >>
    :> 0
/\ count = 0
/\ error = FALSE
/\ handler = ""
/\ history = 0
    :> [bankAfter |->
        <<
            [channel |-> "", id |-> "", port |-> ""], [denom |-> "",
              prefix0 |-> [channel |-> "", port |-> ""],
              prefix1 |-> [channel |-> "", port |-> ""]]
          >>
            :> 0,
      bankBefore |->
        <<
            [channel |-> "", id |-> "", port |-> ""], [denom |-> "",
              prefix0 |-> [channel |-> "", port |-> ""],
              prefix1 |-> [channel |-> "", port |-> ""]]
          >>
            :> 0,
      error |-> FALSE,
      handler |-> "",
      packet |->
        [data |->
            [amount |-> 1,
              denomTrace |->
                [denom |-> "eth",
                  prefix0 |-> [channel |-> "creata-hub", port |-> "creata-hub"],
                  prefix1 |-> [channel |-> "", port |-> ""]],
              receiver |-> "a2",
              sender |-> "a3"],
          destChannel |-> "channel-0",
          destPort |-> "transfer",
          sourceChannel |-> "channel-0",
          sourcePort |-> "transfer"]]
/\ p = [data |->
    [amount |-> 1,
      denomTrace |->
        [denom |-> "eth",
          prefix0 |-> [channel |-> "creata-hub", port |-> "creata-hub"],
          prefix1 |-> [channel |-> "", port |-> ""]],
      receiver |-> "a2",
      sender |-> "a3"],
  destChannel |-> "channel-0",
  destPort |-> "transfer",
  sourceChannel |-> "channel-0",
  sourcePort |-> "transfer"]

(* Transition 2 to State3 *)

State3 ==
/\ bank = <<
    [channel |-> "", id |-> "", port |-> ""], [denom |-> "",
      prefix0 |-> [channel |-> "", port |-> ""],
      prefix1 |-> [channel |-> "", port |-> ""]]
  >>
    :> 0
  @@ <<
    [channel |-> "", id |-> "a2", port |-> ""], [denom |-> "eth",
      prefix0 |-> [channel |-> "creata-hub", port |-> "creata-hub"],
      prefix1 |-> [channel |-> "channel-0", port |-> "transfer"]]
  >>
    :> 1
/\ count = 1
/\ error = FALSE
/\ handler = "OnRecvPacket"
/\ history = 0
    :> [bankAfter |->
        <<
            [channel |-> "", id |-> "", port |-> ""], [denom |-> "",
              prefix0 |-> [channel |-> "", port |-> ""],
              prefix1 |-> [channel |-> "", port |-> ""]]
          >>
            :> 0,
      bankBefore |->
        <<
            [channel |-> "", id |-> "", port |-> ""], [denom |-> "",
              prefix0 |-> [channel |-> "", port |-> ""],
              prefix1 |-> [channel |-> "", port |-> ""]]
          >>
            :> 0,
      error |-> FALSE,
      handler |-> "",
      packet |->
        [data |->
            [amount |-> 1,
              denomTrace |->
                [denom |-> "eth",
                  prefix0 |-> [channel |-> "creata-hub", port |-> "creata-hub"],
                  prefix1 |-> [channel |-> "", port |-> ""]],
              receiver |-> "a2",
              sender |-> "a3"],
          destChannel |-> "channel-0",
          destPort |-> "transfer",
          sourceChannel |-> "channel-0",
          sourcePort |-> "transfer"]]
  @@ 1
    :> [bankAfter |->
        <<
            [channel |-> "", id |-> "", port |-> ""], [denom |-> "",
              prefix0 |-> [channel |-> "", port |-> ""],
              prefix1 |-> [channel |-> "", port |-> ""]]
          >>
            :> 0
          @@ <<
            [channel |-> "", id |-> "a2", port |-> ""], [denom |-> "eth",
              prefix0 |-> [channel |-> "creata-hub", port |-> "creata-hub"],
              prefix1 |-> [channel |-> "channel-0", port |-> "transfer"]]
          >>
            :> 1,
      bankBefore |->
        <<
            [channel |-> "", id |-> "", port |-> ""], [denom |-> "",
              prefix0 |-> [channel |-> "", port |-> ""],
              prefix1 |-> [channel |-> "", port |-> ""]]
          >>
            :> 0,
      error |-> FALSE,
      handler |-> "OnRecvPacket",
      packet |->
        [data |->
            [amount |-> 1,
              denomTrace |->
                [denom |-> "eth",
                  prefix0 |-> [channel |-> "creata-hub", port |-> "creata-hub"],
                  prefix1 |-> [channel |-> "", port |-> ""]],
              receiver |-> "a2",
              sender |-> "a3"],
          destChannel |-> "channel-0",
          destPort |-> "transfer",
          sourceChannel |-> "channel-0",
          sourcePort |-> "transfer"]]
/\ p = [data |->
    [amount |-> 1,
      denomTrace |->
        [denom |-> "eth",
          prefix0 |-> [channel |-> "creata-hub", port |-> "creata-hub"],
          prefix1 |-> [channel |-> "channel-0", port |-> "transfer"]],
      receiver |-> "a1",
      sender |-> "a2"],
  destChannel |-> "channel-0",
  destPort |-> "transfer",
  sourceChannel |-> "channel-1",
  sourcePort |-> "transfer"]

(* Transition 1 to State4 *)

State4 ==
/\ bank = <<
    [channel |-> "", id |-> "", port |-> ""], [denom |-> "",
      prefix0 |-> [channel |-> "", port |-> ""],
      prefix1 |-> [channel |-> "", port |-> ""]]
  >>
    :> 0
  @@ <<
    [channel |-> "", id |-> "a2", port |-> ""], [denom |-> "eth",
      prefix0 |-> [channel |-> "creata-hub", port |-> "creata-hub"],
      prefix1 |-> [channel |-> "channel-0", port |-> "transfer"]]
  >>
    :> 0
  @@ <<
    [channel |-> "channel-1", id |-> "", port |-> "transfer"], [denom |-> "eth",
      prefix0 |-> [channel |-> "creata-hub", port |-> "creata-hub"],
      prefix1 |-> [channel |-> "channel-0", port |-> "transfer"]]
  >>
    :> 1
/\ count = 2
/\ error = FALSE
/\ handler = "SendTransfer"
/\ history = 0
    :> [bankAfter |->
        <<
            [channel |-> "", id |-> "", port |-> ""], [denom |-> "",
              prefix0 |-> [channel |-> "", port |-> ""],
              prefix1 |-> [channel |-> "", port |-> ""]]
          >>
            :> 0,
      bankBefore |->
        <<
            [channel |-> "", id |-> "", port |-> ""], [denom |-> "",
              prefix0 |-> [channel |-> "", port |-> ""],
              prefix1 |-> [channel |-> "", port |-> ""]]
          >>
            :> 0,
      error |-> FALSE,
      handler |-> "",
      packet |->
        [data |->
            [amount |-> 1,
              denomTrace |->
                [denom |-> "eth",
                  prefix0 |-> [channel |-> "creata-hub", port |-> "creata-hub"],
                  prefix1 |-> [channel |-> "", port |-> ""]],
              receiver |-> "a2",
              sender |-> "a3"],
          destChannel |-> "channel-0",
          destPort |-> "transfer",
          sourceChannel |-> "channel-0",
          sourcePort |-> "transfer"]]
  @@ 1
    :> [bankAfter |->
        <<
            [channel |-> "", id |-> "", port |-> ""], [denom |-> "",
              prefix0 |-> [channel |-> "", port |-> ""],
              prefix1 |-> [channel |-> "", port |-> ""]]
          >>
            :> 0
          @@ <<
            [channel |-> "", id |-> "a2", port |-> ""], [denom |-> "eth",
              prefix0 |-> [channel |-> "creata-hub", port |-> "creata-hub"],
              prefix1 |-> [channel |-> "channel-0", port |-> "transfer"]]
          >>
            :> 1,
      bankBefore |->
        <<
            [channel |-> "", id |-> "", port |-> ""], [denom |-> "",
              prefix0 |-> [channel |-> "", port |-> ""],
              prefix1 |-> [channel |-> "", port |-> ""]]
          >>
            :> 0,
      error |-> FALSE,
      handler |-> "OnRecvPacket",
      packet |->
        [data |->
            [amount |-> 1,
              denomTrace |->
                [denom |-> "eth",
                  prefix0 |-> [channel |-> "creata-hub", port |-> "creata-hub"],
                  prefix1 |-> [channel |-> "", port |-> ""]],
              receiver |-> "a2",
              sender |-> "a3"],
          destChannel |-> "channel-0",
          destPort |-> "transfer",
          sourceChannel |-> "channel-0",
          sourcePort |-> "transfer"]]
  @@ 2
    :> [bankAfter |->
        <<
            [channel |-> "", id |-> "", port |-> ""], [denom |-> "",
              prefix0 |-> [channel |-> "", port |-> ""],
              prefix1 |-> [channel |-> "", port |-> ""]]
          >>
            :> 0
          @@ <<
            [channel |-> "", id |-> "a2", port |-> ""], [denom |-> "eth",
              prefix0 |-> [channel |-> "creata-hub", port |-> "creata-hub"],
              prefix1 |-> [channel |-> "channel-0", port |-> "transfer"]]
          >>
            :> 0
          @@ <<
            [channel |-> "channel-1", id |-> "", port |-> "transfer"], [denom |->
                "eth",
              prefix0 |-> [channel |-> "creata-hub", port |-> "creata-hub"],
              prefix1 |-> [channel |-> "channel-0", port |-> "transfer"]]
          >>
            :> 1,
      bankBefore |->
        <<
            [channel |-> "", id |-> "", port |-> ""], [denom |-> "",
              prefix0 |-> [channel |-> "", port |-> ""],
              prefix1 |-> [channel |-> "", port |-> ""]]
          >>
            :> 0
          @@ <<
            [channel |-> "", id |-> "a2", port |-> ""], [denom |-> "eth",
              prefix0 |-> [channel |-> "creata-hub", port |-> "creata-hub"],
              prefix1 |-> [channel |-> "channel-0", port |-> "transfer"]]
          >>
            :> 1,
      error |-> FALSE,
      handler |-> "SendTransfer",
      packet |->
        [data |->
            [amount |-> 1,
              denomTrace |->
                [denom |-> "eth",
                  prefix0 |-> [channel |-> "creata-hub", port |-> "creata-hub"],
                  prefix1 |-> [channel |-> "channel-0", port |-> "transfer"]],
              receiver |-> "a1",
              sender |-> "a2"],
          destChannel |-> "channel-0",
          destPort |-> "transfer",
          sourceChannel |-> "channel-1",
          sourcePort |-> "transfer"]]
/\ p = [data |->
    [amount |-> 0,
      denomTrace |->
        [denom |-> "",
          prefix0 |-> [channel |-> "", port |-> ""],
          prefix1 |-> [channel |-> "", port |-> ""]],
      receiver |-> "",
      sender |-> ""],
  destChannel |-> "",
  destPort |-> "",
  sourceChannel |-> "",
  sourcePort |-> ""]

(* The following formula holds true in the last state and violates the invariant *)

InvariantViolation ==
  BMC!Skolem((\E s$2 \in DOMAIN history:
    history[s$2]["handler"] = "SendTransfer"
      /\ history[s$2]["error"] = FALSE
      /\ history[s$2]["packet"]["data"]["amount"] > 0))

================================================================================
\* Created by Apalache on Thu Dec 10 10:58:54 CET 2020
\* https://github.com/informalsystems/apalache
