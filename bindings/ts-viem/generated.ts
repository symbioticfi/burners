//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// IAddressRequests
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

export const iAddressRequestsAbi = [
  {
    type: "function",
    inputs: [
      { name: "index", internalType: "uint256", type: "uint256" },
      { name: "maxRequestIds", internalType: "uint256", type: "uint256" },
    ],
    name: "requestIds",
    outputs: [{ name: "requestIds", internalType: "address[]", type: "address[]" }],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [],
    name: "requestIdsLength",
    outputs: [{ name: "", internalType: "uint256", type: "uint256" }],
    stateMutability: "view",
  },
  { type: "error", inputs: [], name: "InvalidRequestId" },
] as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// IBurnerRouter
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

export const iBurnerRouterAbi = [
  { type: "function", inputs: [], name: "acceptDelay", outputs: [], stateMutability: "nonpayable" },
  { type: "function", inputs: [], name: "acceptGlobalReceiver", outputs: [], stateMutability: "nonpayable" },
  {
    type: "function",
    inputs: [{ name: "network", internalType: "address", type: "address" }],
    name: "acceptNetworkReceiver",
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    inputs: [
      { name: "network", internalType: "address", type: "address" },
      { name: "operator", internalType: "address", type: "address" },
    ],
    name: "acceptOperatorNetworkReceiver",
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    inputs: [{ name: "receiver", internalType: "address", type: "address" }],
    name: "balanceOf",
    outputs: [{ name: "", internalType: "uint256", type: "uint256" }],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [],
    name: "collateral",
    outputs: [{ name: "", internalType: "address", type: "address" }],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [],
    name: "delay",
    outputs: [{ name: "", internalType: "uint48", type: "uint48" }],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [],
    name: "globalReceiver",
    outputs: [{ name: "", internalType: "address", type: "address" }],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [],
    name: "lastBalance",
    outputs: [{ name: "", internalType: "uint256", type: "uint256" }],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [{ name: "network", internalType: "address", type: "address" }],
    name: "networkReceiver",
    outputs: [{ name: "", internalType: "address", type: "address" }],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [
      { name: "subnetwork", internalType: "bytes32", type: "bytes32" },
      { name: "operator", internalType: "address", type: "address" },
      { name: "amount", internalType: "uint256", type: "uint256" },
      { name: "captureTimestamp", internalType: "uint48", type: "uint48" },
    ],
    name: "onSlash",
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    inputs: [
      { name: "network", internalType: "address", type: "address" },
      { name: "operator", internalType: "address", type: "address" },
    ],
    name: "operatorNetworkReceiver",
    outputs: [{ name: "", internalType: "address", type: "address" }],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [],
    name: "pendingDelay",
    outputs: [
      { name: "", internalType: "uint48", type: "uint48" },
      { name: "", internalType: "uint48", type: "uint48" },
    ],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [],
    name: "pendingGlobalReceiver",
    outputs: [
      { name: "", internalType: "address", type: "address" },
      { name: "", internalType: "uint48", type: "uint48" },
    ],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [{ name: "network", internalType: "address", type: "address" }],
    name: "pendingNetworkReceiver",
    outputs: [
      { name: "", internalType: "address", type: "address" },
      { name: "", internalType: "uint48", type: "uint48" },
    ],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [
      { name: "network", internalType: "address", type: "address" },
      { name: "operator", internalType: "address", type: "address" },
    ],
    name: "pendingOperatorNetworkReceiver",
    outputs: [
      { name: "", internalType: "address", type: "address" },
      { name: "", internalType: "uint48", type: "uint48" },
    ],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [{ name: "newDelay", internalType: "uint48", type: "uint48" }],
    name: "setDelay",
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    inputs: [{ name: "receiver", internalType: "address", type: "address" }],
    name: "setGlobalReceiver",
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    inputs: [
      { name: "network", internalType: "address", type: "address" },
      { name: "receiver", internalType: "address", type: "address" },
    ],
    name: "setNetworkReceiver",
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    inputs: [
      { name: "network", internalType: "address", type: "address" },
      { name: "operator", internalType: "address", type: "address" },
      { name: "receiver", internalType: "address", type: "address" },
    ],
    name: "setOperatorNetworkReceiver",
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    inputs: [{ name: "receiver", internalType: "address", type: "address" }],
    name: "triggerTransfer",
    outputs: [{ name: "amount", internalType: "uint256", type: "uint256" }],
    stateMutability: "nonpayable",
  },
  { type: "event", anonymous: false, inputs: [], name: "AcceptDelay" },
  { type: "event", anonymous: false, inputs: [], name: "AcceptGlobalReceiver" },
  {
    type: "event",
    anonymous: false,
    inputs: [{ name: "network", internalType: "address", type: "address", indexed: true }],
    name: "AcceptNetworkReceiver",
  },
  {
    type: "event",
    anonymous: false,
    inputs: [
      { name: "network", internalType: "address", type: "address", indexed: true },
      { name: "operator", internalType: "address", type: "address", indexed: true },
    ],
    name: "AcceptOperatorNetworkReceiver",
  },
  {
    type: "event",
    anonymous: false,
    inputs: [{ name: "delay", internalType: "uint48", type: "uint48", indexed: false }],
    name: "SetDelay",
  },
  {
    type: "event",
    anonymous: false,
    inputs: [{ name: "receiver", internalType: "address", type: "address", indexed: false }],
    name: "SetGlobalReceiver",
  },
  {
    type: "event",
    anonymous: false,
    inputs: [
      { name: "network", internalType: "address", type: "address", indexed: true },
      { name: "receiver", internalType: "address", type: "address", indexed: false },
    ],
    name: "SetNetworkReceiver",
  },
  {
    type: "event",
    anonymous: false,
    inputs: [
      { name: "network", internalType: "address", type: "address", indexed: true },
      { name: "operator", internalType: "address", type: "address", indexed: true },
      { name: "receiver", internalType: "address", type: "address", indexed: false },
    ],
    name: "SetOperatorNetworkReceiver",
  },
  {
    type: "event",
    anonymous: false,
    inputs: [
      { name: "receiver", internalType: "address", type: "address", indexed: true },
      { name: "amount", internalType: "uint256", type: "uint256", indexed: false },
    ],
    name: "TriggerTransfer",
  },
  { type: "error", inputs: [], name: "AlreadySet" },
  { type: "error", inputs: [], name: "DuplicateNetworkReceiver" },
  { type: "error", inputs: [], name: "DuplicateOperatorNetworkReceiver" },
  { type: "error", inputs: [], name: "InsufficientBalance" },
  { type: "error", inputs: [], name: "InvalidCollateral" },
  { type: "error", inputs: [], name: "InvalidReceiver" },
  { type: "error", inputs: [], name: "InvalidReceiverSetEpochsDelay" },
  { type: "error", inputs: [], name: "NotReady" },
] as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// IBurnerRouterFactory
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

export const iBurnerRouterFactoryAbi = [
  {
    type: "function",
    inputs: [
      {
        name: "params",
        internalType: "struct IBurnerRouter.InitParams",
        type: "tuple",
        components: [
          { name: "owner", internalType: "address", type: "address" },
          { name: "collateral", internalType: "address", type: "address" },
          { name: "delay", internalType: "uint48", type: "uint48" },
          { name: "globalReceiver", internalType: "address", type: "address" },
          {
            name: "networkReceivers",
            internalType: "struct IBurnerRouter.NetworkReceiver[]",
            type: "tuple[]",
            components: [
              { name: "network", internalType: "address", type: "address" },
              { name: "receiver", internalType: "address", type: "address" },
            ],
          },
          {
            name: "operatorNetworkReceivers",
            internalType: "struct IBurnerRouter.OperatorNetworkReceiver[]",
            type: "tuple[]",
            components: [
              { name: "network", internalType: "address", type: "address" },
              { name: "operator", internalType: "address", type: "address" },
              { name: "receiver", internalType: "address", type: "address" },
            ],
          },
        ],
      },
    ],
    name: "create",
    outputs: [{ name: "", internalType: "address", type: "address" }],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    inputs: [{ name: "index", internalType: "uint256", type: "uint256" }],
    name: "entity",
    outputs: [{ name: "", internalType: "address", type: "address" }],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [{ name: "account", internalType: "address", type: "address" }],
    name: "isEntity",
    outputs: [{ name: "", internalType: "bool", type: "bool" }],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [],
    name: "totalEntities",
    outputs: [{ name: "", internalType: "uint256", type: "uint256" }],
    stateMutability: "view",
  },
  {
    type: "event",
    anonymous: false,
    inputs: [{ name: "entity", internalType: "address", type: "address", indexed: true }],
    name: "AddEntity",
  },
  { type: "error", inputs: [], name: "EntityNotExist" },
] as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// IETHx_Burner
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

export const ietHxBurnerAbi = [
  {
    type: "function",
    inputs: [],
    name: "COLLATERAL",
    outputs: [{ name: "", internalType: "address", type: "address" }],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [],
    name: "STADER_CONFIG",
    outputs: [{ name: "", internalType: "address", type: "address" }],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [],
    name: "STAKE_POOLS_MANAGER",
    outputs: [{ name: "", internalType: "address", type: "address" }],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [],
    name: "USER_WITHDRAW_MANAGER",
    outputs: [{ name: "", internalType: "address", type: "address" }],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [
      { name: "index", internalType: "uint256", type: "uint256" },
      { name: "maxRequestIds", internalType: "uint256", type: "uint256" },
    ],
    name: "requestIds",
    outputs: [{ name: "requestIds", internalType: "uint256[]", type: "uint256[]" }],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [],
    name: "requestIdsLength",
    outputs: [{ name: "", internalType: "uint256", type: "uint256" }],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [{ name: "requestId", internalType: "uint256", type: "uint256" }],
    name: "triggerBurn",
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    inputs: [{ name: "maxWithdrawalAmount", internalType: "uint256", type: "uint256" }],
    name: "triggerWithdrawal",
    outputs: [{ name: "requestId", internalType: "uint256", type: "uint256" }],
    stateMutability: "nonpayable",
  },
  {
    type: "event",
    anonymous: false,
    inputs: [
      { name: "caller", internalType: "address", type: "address", indexed: true },
      { name: "requestId", internalType: "uint256", type: "uint256", indexed: false },
    ],
    name: "TriggerBurn",
  },
  {
    type: "event",
    anonymous: false,
    inputs: [
      { name: "caller", internalType: "address", type: "address", indexed: true },
      { name: "requestId", internalType: "uint256", type: "uint256", indexed: false },
    ],
    name: "TriggerWithdrawal",
  },
  { type: "error", inputs: [], name: "InvalidETHxMaximumWithdrawal" },
  { type: "error", inputs: [], name: "InvalidRequestId" },
] as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// IEthenaMinting
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

export const iEthenaMintingAbi = [
  {
    type: "function",
    inputs: [{ name: "benefactor", internalType: "address", type: "address" }],
    name: "addWhitelistedBenefactor",
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    inputs: [],
    name: "globalConfig",
    outputs: [
      {
        name: "",
        internalType: "struct IEthenaMinting.GlobalConfig",
        type: "tuple",
        components: [
          { name: "globalMaxMintPerBlock", internalType: "uint128", type: "uint128" },
          { name: "globalMaxRedeemPerBlock", internalType: "uint128", type: "uint128" },
        ],
      },
    ],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [
      {
        name: "order",
        internalType: "struct IEthenaMinting.Order",
        type: "tuple",
        components: [
          { name: "order_id", internalType: "string", type: "string" },
          { name: "order_type", internalType: "enum IEthenaMinting.OrderType", type: "uint8" },
          { name: "expiry", internalType: "uint120", type: "uint120" },
          { name: "nonce", internalType: "uint128", type: "uint128" },
          { name: "benefactor", internalType: "address", type: "address" },
          { name: "beneficiary", internalType: "address", type: "address" },
          { name: "collateral_asset", internalType: "address", type: "address" },
          { name: "collateral_amount", internalType: "uint128", type: "uint128" },
          { name: "usde_amount", internalType: "uint128", type: "uint128" },
        ],
      },
    ],
    name: "hashOrder",
    outputs: [{ name: "", internalType: "bytes32", type: "bytes32" }],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [
      {
        name: "order",
        internalType: "struct IEthenaMinting.Order",
        type: "tuple",
        components: [
          { name: "order_id", internalType: "string", type: "string" },
          { name: "order_type", internalType: "enum IEthenaMinting.OrderType", type: "uint8" },
          { name: "expiry", internalType: "uint120", type: "uint120" },
          { name: "nonce", internalType: "uint128", type: "uint128" },
          { name: "benefactor", internalType: "address", type: "address" },
          { name: "beneficiary", internalType: "address", type: "address" },
          { name: "collateral_asset", internalType: "address", type: "address" },
          { name: "collateral_amount", internalType: "uint128", type: "uint128" },
          { name: "usde_amount", internalType: "uint128", type: "uint128" },
        ],
      },
      {
        name: "signature",
        internalType: "struct IEthenaMinting.Signature",
        type: "tuple",
        components: [
          { name: "signature_type", internalType: "enum IEthenaMinting.SignatureType", type: "uint8" },
          { name: "signature_bytes", internalType: "bytes", type: "bytes" },
        ],
      },
    ],
    name: "redeem",
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    inputs: [{ name: "asset", internalType: "address", type: "address" }],
    name: "tokenConfig",
    outputs: [
      {
        name: "",
        internalType: "struct IEthenaMinting.TokenConfig",
        type: "tuple",
        components: [
          { name: "tokenType", internalType: "enum IEthenaMinting.TokenType", type: "uint8" },
          { name: "isActive", internalType: "bool", type: "bool" },
          { name: "maxMintPerBlock", internalType: "uint128", type: "uint128" },
          { name: "maxRedeemPerBlock", internalType: "uint128", type: "uint128" },
        ],
      },
    ],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [{ name: "blockNumber", internalType: "uint256", type: "uint256" }],
    name: "totalPerBlock",
    outputs: [
      {
        name: "",
        internalType: "struct IEthenaMinting.BlockTotals",
        type: "tuple",
        components: [
          { name: "mintedPerBlock", internalType: "uint128", type: "uint128" },
          { name: "redeemedPerBlock", internalType: "uint128", type: "uint128" },
        ],
      },
    ],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [
      { name: "blockNumber", internalType: "uint256", type: "uint256" },
      { name: "asset", internalType: "address", type: "address" },
    ],
    name: "totalPerBlockPerAsset",
    outputs: [
      {
        name: "",
        internalType: "struct IEthenaMinting.BlockTotals",
        type: "tuple",
        components: [
          { name: "mintedPerBlock", internalType: "uint128", type: "uint128" },
          { name: "redeemedPerBlock", internalType: "uint128", type: "uint128" },
        ],
      },
    ],
    stateMutability: "view",
  },
] as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// IFraxEtherRedemptionQueue
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

export const iFraxEtherRedemptionQueueAbi = [
  {
    type: "function",
    inputs: [
      { name: "_nftId", internalType: "uint256", type: "uint256" },
      { name: "_recipient", internalType: "address payable", type: "address" },
    ],
    name: "burnRedemptionTicketNft",
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    inputs: [
      { name: "_recipient", internalType: "address", type: "address" },
      { name: "_sfrxEthAmount", internalType: "uint120", type: "uint120" },
    ],
    name: "enterRedemptionQueueViaSfrxEth",
    outputs: [{ name: "_nftId", internalType: "uint256", type: "uint256" }],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    inputs: [],
    name: "redemptionQueueState",
    outputs: [
      { name: "nextNftId", internalType: "uint64", type: "uint64" },
      { name: "queueLengthSecs", internalType: "uint64", type: "uint64" },
      { name: "redemptionFee", internalType: "uint64", type: "uint64" },
      { name: "earlyExitFee", internalType: "uint64", type: "uint64" },
    ],
    stateMutability: "view",
  },
] as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// IMETH
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

export const imethAbi = [
  {
    type: "function",
    inputs: [
      { name: "staker", internalType: "address", type: "address" },
      { name: "amount", internalType: "uint256", type: "uint256" },
    ],
    name: "mint",
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    inputs: [],
    name: "stakingContract",
    outputs: [{ name: "", internalType: "address", type: "address" }],
    stateMutability: "view",
  },
] as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// IRocketTokenRETH
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

export const iRocketTokenRethAbi = [
  {
    type: "function",
    inputs: [{ name: "_rethAmount", internalType: "uint256", type: "uint256" }],
    name: "burn",
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    inputs: [{ name: "_rethAmount", internalType: "uint256", type: "uint256" }],
    name: "getEthValue",
    outputs: [{ name: "", internalType: "uint256", type: "uint256" }],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [{ name: "_ethAmount", internalType: "uint256", type: "uint256" }],
    name: "getRethValue",
    outputs: [{ name: "", internalType: "uint256", type: "uint256" }],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [],
    name: "getTotalCollateral",
    outputs: [{ name: "", internalType: "uint256", type: "uint256" }],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [
      { name: "_ethAmount", internalType: "uint256", type: "uint256" },
      { name: "_to", internalType: "address", type: "address" },
    ],
    name: "mint",
    outputs: [],
    stateMutability: "nonpayable",
  },
] as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// ISUSDe
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

export const isusDeAbi = [
  {
    type: "function",
    inputs: [],
    name: "asset",
    outputs: [{ name: "", internalType: "address", type: "address" }],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [],
    name: "cooldownDuration",
    outputs: [{ name: "", internalType: "uint24", type: "uint24" }],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [{ name: "shares", internalType: "uint256", type: "uint256" }],
    name: "cooldownShares",
    outputs: [{ name: "assets", internalType: "uint256", type: "uint256" }],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    inputs: [
      { name: "assets", internalType: "uint256", type: "uint256" },
      { name: "receiver", internalType: "address", type: "address" },
    ],
    name: "deposit",
    outputs: [{ name: "", internalType: "uint256", type: "uint256" }],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    inputs: [{ name: "shares", internalType: "uint256", type: "uint256" }],
    name: "previewRedeem",
    outputs: [{ name: "", internalType: "uint256", type: "uint256" }],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    inputs: [
      { name: "shares", internalType: "uint256", type: "uint256" },
      { name: "receiver", internalType: "address", type: "address" },
      { name: "_owner", internalType: "address", type: "address" },
    ],
    name: "redeem",
    outputs: [{ name: "", internalType: "uint256", type: "uint256" }],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    inputs: [{ name: "duration", internalType: "uint24", type: "uint24" }],
    name: "setCooldownDuration",
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    inputs: [{ name: "receiver", internalType: "address", type: "address" }],
    name: "unstake",
    outputs: [],
    stateMutability: "nonpayable",
  },
] as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// IStaderConfig
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

export const iStaderConfigAbi = [
  {
    type: "function",
    inputs: [],
    name: "getMaxWithdrawAmount",
    outputs: [{ name: "", internalType: "uint256", type: "uint256" }],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [],
    name: "getMinBlockDelayToFinalizeWithdrawRequest",
    outputs: [{ name: "", internalType: "uint256", type: "uint256" }],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [],
    name: "getMinWithdrawAmount",
    outputs: [{ name: "", internalType: "uint256", type: "uint256" }],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [],
    name: "getStakePoolManager",
    outputs: [{ name: "", internalType: "address", type: "address" }],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [],
    name: "getUserWithdrawManager",
    outputs: [{ name: "", internalType: "address", type: "address" }],
    stateMutability: "view",
  },
] as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// IStaderStakePoolsManager
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

export const iStaderStakePoolsManagerAbi = [
  {
    type: "function",
    inputs: [{ name: "_assets", internalType: "uint256", type: "uint256" }],
    name: "previewDeposit",
    outputs: [{ name: "", internalType: "uint256", type: "uint256" }],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [{ name: "_shares", internalType: "uint256", type: "uint256" }],
    name: "previewWithdraw",
    outputs: [{ name: "", internalType: "uint256", type: "uint256" }],
    stateMutability: "view",
  },
] as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// IStaking
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

export const iStakingAbi = [
  {
    type: "function",
    inputs: [{ name: "unstakeRequestID", internalType: "uint256", type: "uint256" }],
    name: "claimUnstakeRequest",
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    inputs: [{ name: "mETHAmount", internalType: "uint256", type: "uint256" }],
    name: "mETHToETH",
    outputs: [{ name: "", internalType: "uint256", type: "uint256" }],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [],
    name: "minimumUnstakeBound",
    outputs: [{ name: "", internalType: "uint256", type: "uint256" }],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [
      { name: "methAmount", internalType: "uint128", type: "uint128" },
      { name: "minETHAmount", internalType: "uint128", type: "uint128" },
    ],
    name: "unstakeRequest",
    outputs: [{ name: "", internalType: "uint256", type: "uint256" }],
    stateMutability: "nonpayable",
  },
] as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// ISwETH
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

export const iSwEthAbi = [
  { type: "function", inputs: [], name: "deposit", outputs: [], stateMutability: "payable" },
] as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// ISwEXIT
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

export const iSwExitAbi = [
  {
    type: "function",
    inputs: [{ name: "amount", internalType: "uint256", type: "uint256" }],
    name: "createWithdrawRequest",
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    inputs: [{ name: "tokenId", internalType: "uint256", type: "uint256" }],
    name: "finalizeWithdrawal",
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    inputs: [],
    name: "getLastTokenIdCreated",
    outputs: [{ name: "", internalType: "uint256", type: "uint256" }],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [{ name: "_lastTokenIdToProcess", internalType: "uint256", type: "uint256" }],
    name: "processWithdrawals",
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    inputs: [],
    name: "withdrawRequestMaximum",
    outputs: [{ name: "", internalType: "uint256", type: "uint256" }],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [],
    name: "withdrawRequestMinimum",
    outputs: [{ name: "", internalType: "uint256", type: "uint256" }],
    stateMutability: "view",
  },
] as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// IUSDe
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

export const iusDeAbi = [
  {
    type: "function",
    inputs: [{ name: "value", internalType: "uint256", type: "uint256" }],
    name: "burn",
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    inputs: [
      { name: "account", internalType: "address", type: "address" },
      { name: "amount", internalType: "uint256", type: "uint256" },
    ],
    name: "mint",
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    inputs: [],
    name: "minter",
    outputs: [{ name: "", internalType: "address", type: "address" }],
    stateMutability: "view",
  },
] as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// IUintRequests
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

export const iUintRequestsAbi = [
  {
    type: "function",
    inputs: [
      { name: "index", internalType: "uint256", type: "uint256" },
      { name: "maxRequestIds", internalType: "uint256", type: "uint256" },
    ],
    name: "requestIds",
    outputs: [{ name: "requestIds", internalType: "uint256[]", type: "uint256[]" }],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [],
    name: "requestIdsLength",
    outputs: [{ name: "", internalType: "uint256", type: "uint256" }],
    stateMutability: "view",
  },
  { type: "error", inputs: [], name: "InvalidRequestId" },
] as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// IUserWithdrawalManager
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

export const iUserWithdrawalManagerAbi = [
  {
    type: "function",
    inputs: [{ name: "_requestId", internalType: "uint256", type: "uint256" }],
    name: "claim",
    outputs: [],
    stateMutability: "nonpayable",
  },
  { type: "function", inputs: [], name: "finalizeUserWithdrawalRequest", outputs: [], stateMutability: "nonpayable" },
  {
    type: "function",
    inputs: [],
    name: "nextRequestId",
    outputs: [{ name: "", internalType: "uint256", type: "uint256" }],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [],
    name: "nextRequestIdToFinalize",
    outputs: [{ name: "", internalType: "uint256", type: "uint256" }],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [
      { name: "_ethXAmount", internalType: "uint256", type: "uint256" },
      { name: "_owner", internalType: "address", type: "address" },
    ],
    name: "requestWithdraw",
    outputs: [{ name: "", internalType: "uint256", type: "uint256" }],
    stateMutability: "nonpayable",
  },
] as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// IWithdrawalQueue
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

export const iWithdrawalQueueAbi = [
  {
    type: "function",
    inputs: [],
    name: "MAX_STETH_WITHDRAWAL_AMOUNT",
    outputs: [{ name: "", internalType: "uint256", type: "uint256" }],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [],
    name: "MIN_STETH_WITHDRAWAL_AMOUNT",
    outputs: [{ name: "", internalType: "uint256", type: "uint256" }],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [],
    name: "STETH",
    outputs: [{ name: "", internalType: "address", type: "address" }],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [{ name: "_requestId", internalType: "uint256", type: "uint256" }],
    name: "claimWithdrawal",
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    inputs: [
      { name: "_requestIds", internalType: "uint256[]", type: "uint256[]" },
      { name: "_hints", internalType: "uint256[]", type: "uint256[]" },
    ],
    name: "claimWithdrawals",
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    inputs: [
      { name: "_lastRequestIdToBeFinalized", internalType: "uint256", type: "uint256" },
      { name: "_maxShareRate", internalType: "uint256", type: "uint256" },
    ],
    name: "finalize",
    outputs: [],
    stateMutability: "payable",
  },
  {
    type: "function",
    inputs: [
      { name: "_requestIds", internalType: "uint256[]", type: "uint256[]" },
      { name: "_firstIndex", internalType: "uint256", type: "uint256" },
      { name: "_lastIndex", internalType: "uint256", type: "uint256" },
    ],
    name: "findCheckpointHints",
    outputs: [{ name: "hintIds", internalType: "uint256[]", type: "uint256[]" }],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [],
    name: "getLastCheckpointIndex",
    outputs: [{ name: "", internalType: "uint256", type: "uint256" }],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [],
    name: "getLastRequestId",
    outputs: [{ name: "", internalType: "uint256", type: "uint256" }],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [
      { name: "_amounts", internalType: "uint256[]", type: "uint256[]" },
      { name: "_owner", internalType: "address", type: "address" },
    ],
    name: "requestWithdrawals",
    outputs: [{ name: "requestIds", internalType: "uint256[]", type: "uint256[]" }],
    stateMutability: "nonpayable",
  },
] as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// IWstETH
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

export const iWstEthAbi = [
  {
    type: "function",
    inputs: [{ name: "_wstETHAmount", internalType: "uint256", type: "uint256" }],
    name: "getStETHByWstETH",
    outputs: [{ name: "", internalType: "uint256", type: "uint256" }],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    inputs: [{ name: "_wstETHAmount", internalType: "uint256", type: "uint256" }],
    name: "unwrap",
    outputs: [{ name: "", internalType: "uint256", type: "uint256" }],
    stateMutability: "nonpayable",
  },
] as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// ImETH_Burner
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

export const imEthBurnerAbi = [
  {
    type: "function",
    inputs: [],
    name: "COLLATERAL",
    outputs: [{ name: "", internalType: "address", type: "address" }],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [],
    name: "STAKING",
    outputs: [{ name: "", internalType: "address", type: "address" }],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [
      { name: "index", internalType: "uint256", type: "uint256" },
      { name: "maxRequestIds", internalType: "uint256", type: "uint256" },
    ],
    name: "requestIds",
    outputs: [{ name: "requestIds", internalType: "uint256[]", type: "uint256[]" }],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [],
    name: "requestIdsLength",
    outputs: [{ name: "", internalType: "uint256", type: "uint256" }],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [{ name: "requestId", internalType: "uint256", type: "uint256" }],
    name: "triggerBurn",
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    inputs: [],
    name: "triggerWithdrawal",
    outputs: [{ name: "requestId", internalType: "uint256", type: "uint256" }],
    stateMutability: "nonpayable",
  },
  {
    type: "event",
    anonymous: false,
    inputs: [
      { name: "caller", internalType: "address", type: "address", indexed: true },
      { name: "requestId", internalType: "uint256", type: "uint256", indexed: false },
    ],
    name: "TriggerBurn",
  },
  {
    type: "event",
    anonymous: false,
    inputs: [
      { name: "caller", internalType: "address", type: "address", indexed: true },
      { name: "requestId", internalType: "uint256", type: "uint256", indexed: false },
    ],
    name: "TriggerWithdrawal",
  },
  { type: "error", inputs: [], name: "InvalidRequestId" },
] as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// IrETH_Burner
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

export const irEthBurnerAbi = [
  {
    type: "function",
    inputs: [],
    name: "COLLATERAL",
    outputs: [{ name: "", internalType: "address", type: "address" }],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [{ name: "amount", internalType: "uint256", type: "uint256" }],
    name: "triggerBurn",
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "event",
    anonymous: false,
    inputs: [
      { name: "caller", internalType: "address", type: "address", indexed: true },
      { name: "assetAmount", internalType: "uint256", type: "uint256", indexed: false },
      { name: "ethAmount", internalType: "uint256", type: "uint256", indexed: false },
    ],
    name: "TriggerBurn",
  },
] as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// IsUSDe_Burner
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

export const isUsDeBurnerAbi = [
  {
    type: "function",
    inputs: [],
    name: "COLLATERAL",
    outputs: [{ name: "", internalType: "address", type: "address" }],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [],
    name: "USDE",
    outputs: [{ name: "", internalType: "address", type: "address" }],
    stateMutability: "view",
  },
  { type: "function", inputs: [], name: "approveUSDeMinter", outputs: [], stateMutability: "nonpayable" },
  {
    type: "function",
    inputs: [
      { name: "index", internalType: "uint256", type: "uint256" },
      { name: "maxRequestIds", internalType: "uint256", type: "uint256" },
    ],
    name: "requestIds",
    outputs: [{ name: "requestIds", internalType: "address[]", type: "address[]" }],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [],
    name: "requestIdsLength",
    outputs: [{ name: "", internalType: "uint256", type: "uint256" }],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [{ name: "asset", internalType: "address", type: "address" }],
    name: "triggerBurn",
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    inputs: [{ name: "requestId", internalType: "address", type: "address" }],
    name: "triggerClaim",
    outputs: [],
    stateMutability: "nonpayable",
  },
  { type: "function", inputs: [], name: "triggerInstantClaim", outputs: [], stateMutability: "nonpayable" },
  {
    type: "function",
    inputs: [],
    name: "triggerWithdrawal",
    outputs: [{ name: "requestId", internalType: "address", type: "address" }],
    stateMutability: "nonpayable",
  },
  {
    type: "event",
    anonymous: false,
    inputs: [
      { name: "caller", internalType: "address", type: "address", indexed: true },
      { name: "asset", internalType: "address", type: "address", indexed: true },
      { name: "amount", internalType: "uint256", type: "uint256", indexed: false },
    ],
    name: "TriggerBurn",
  },
  {
    type: "event",
    anonymous: false,
    inputs: [
      { name: "caller", internalType: "address", type: "address", indexed: true },
      { name: "requestId", internalType: "address", type: "address", indexed: false },
    ],
    name: "TriggerClaim",
  },
  {
    type: "event",
    anonymous: false,
    inputs: [
      { name: "caller", internalType: "address", type: "address", indexed: true },
      { name: "amount", internalType: "uint256", type: "uint256", indexed: false },
    ],
    name: "TriggerInstantClaim",
  },
  {
    type: "event",
    anonymous: false,
    inputs: [
      { name: "caller", internalType: "address", type: "address", indexed: true },
      { name: "amount", internalType: "uint256", type: "uint256", indexed: false },
      { name: "requestId", internalType: "address", type: "address", indexed: false },
    ],
    name: "TriggerWithdrawal",
  },
  { type: "error", inputs: [], name: "HasCooldown" },
  { type: "error", inputs: [], name: "InvalidAsset" },
  { type: "error", inputs: [], name: "InvalidRequestId" },
  { type: "error", inputs: [], name: "NoCooldown" },
  { type: "error", inputs: [], name: "SufficientApproval" },
] as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// IsfrxETH_Burner
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

export const isfrxEthBurnerAbi = [
  {
    type: "function",
    inputs: [],
    name: "COLLATERAL",
    outputs: [{ name: "", internalType: "address", type: "address" }],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [],
    name: "FRAX_ETHER_REDEMPTION_QUEUE",
    outputs: [{ name: "", internalType: "address", type: "address" }],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [
      { name: "index", internalType: "uint256", type: "uint256" },
      { name: "maxRequestIds", internalType: "uint256", type: "uint256" },
    ],
    name: "requestIds",
    outputs: [{ name: "requestIds", internalType: "uint256[]", type: "uint256[]" }],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [],
    name: "requestIdsLength",
    outputs: [{ name: "", internalType: "uint256", type: "uint256" }],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [{ name: "requestId", internalType: "uint256", type: "uint256" }],
    name: "triggerBurn",
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    inputs: [],
    name: "triggerWithdrawal",
    outputs: [{ name: "requestId", internalType: "uint256", type: "uint256" }],
    stateMutability: "nonpayable",
  },
  {
    type: "event",
    anonymous: false,
    inputs: [
      { name: "caller", internalType: "address", type: "address", indexed: true },
      { name: "requestId", internalType: "uint256", type: "uint256", indexed: false },
    ],
    name: "TriggerBurn",
  },
  {
    type: "event",
    anonymous: false,
    inputs: [
      { name: "caller", internalType: "address", type: "address", indexed: true },
      { name: "requestId", internalType: "uint256", type: "uint256", indexed: false },
    ],
    name: "TriggerWithdrawal",
  },
  { type: "error", inputs: [], name: "InvalidRequestId" },
] as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// IswETH_Burner
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

export const iswEthBurnerAbi = [
  {
    type: "function",
    inputs: [],
    name: "COLLATERAL",
    outputs: [{ name: "", internalType: "address", type: "address" }],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [],
    name: "SWEXIT",
    outputs: [{ name: "", internalType: "address", type: "address" }],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [
      { name: "index", internalType: "uint256", type: "uint256" },
      { name: "maxRequestIds", internalType: "uint256", type: "uint256" },
    ],
    name: "requestIds",
    outputs: [{ name: "requestIds", internalType: "uint256[]", type: "uint256[]" }],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [],
    name: "requestIdsLength",
    outputs: [{ name: "", internalType: "uint256", type: "uint256" }],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [{ name: "requestId", internalType: "uint256", type: "uint256" }],
    name: "triggerBurn",
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    inputs: [{ name: "maxRequests", internalType: "uint256", type: "uint256" }],
    name: "triggerWithdrawal",
    outputs: [
      { name: "firstRequestId", internalType: "uint256", type: "uint256" },
      { name: "lastRequestId", internalType: "uint256", type: "uint256" },
    ],
    stateMutability: "nonpayable",
  },
  {
    type: "event",
    anonymous: false,
    inputs: [
      { name: "caller", internalType: "address", type: "address", indexed: true },
      { name: "requestId", internalType: "uint256", type: "uint256", indexed: false },
    ],
    name: "TriggerBurn",
  },
  {
    type: "event",
    anonymous: false,
    inputs: [
      { name: "caller", internalType: "address", type: "address", indexed: true },
      { name: "firstRequestId", internalType: "uint256", type: "uint256", indexed: false },
      { name: "lastRequestId", internalType: "uint256", type: "uint256", indexed: false },
    ],
    name: "TriggerWithdrawal",
  },
  { type: "error", inputs: [], name: "InsufficientWithdrawal" },
  { type: "error", inputs: [], name: "InvalidRequestId" },
] as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// IwstETH_Burner
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

export const iwstEthBurnerAbi = [
  {
    type: "function",
    inputs: [],
    name: "COLLATERAL",
    outputs: [{ name: "", internalType: "address", type: "address" }],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [],
    name: "LIDO_WITHDRAWAL_QUEUE",
    outputs: [{ name: "", internalType: "address", type: "address" }],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [],
    name: "MAX_STETH_WITHDRAWAL_AMOUNT",
    outputs: [{ name: "", internalType: "uint256", type: "uint256" }],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [],
    name: "MIN_STETH_WITHDRAWAL_AMOUNT",
    outputs: [{ name: "", internalType: "uint256", type: "uint256" }],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [],
    name: "STETH",
    outputs: [{ name: "", internalType: "address", type: "address" }],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [
      { name: "index", internalType: "uint256", type: "uint256" },
      { name: "maxRequestIds", internalType: "uint256", type: "uint256" },
    ],
    name: "requestIds",
    outputs: [{ name: "requestIds", internalType: "uint256[]", type: "uint256[]" }],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [],
    name: "requestIdsLength",
    outputs: [{ name: "", internalType: "uint256", type: "uint256" }],
    stateMutability: "view",
  },
  {
    type: "function",
    inputs: [{ name: "requestId", internalType: "uint256", type: "uint256" }],
    name: "triggerBurn",
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    inputs: [
      { name: "requestIds", internalType: "uint256[]", type: "uint256[]" },
      { name: "hints", internalType: "uint256[]", type: "uint256[]" },
    ],
    name: "triggerBurnBatch",
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    inputs: [{ name: "maxRequests", internalType: "uint256", type: "uint256" }],
    name: "triggerWithdrawal",
    outputs: [{ name: "requestIds", internalType: "uint256[]", type: "uint256[]" }],
    stateMutability: "nonpayable",
  },
  {
    type: "event",
    anonymous: false,
    inputs: [
      { name: "caller", internalType: "address", type: "address", indexed: true },
      { name: "requestId", internalType: "uint256", type: "uint256", indexed: false },
    ],
    name: "TriggerBurn",
  },
  {
    type: "event",
    anonymous: false,
    inputs: [
      { name: "caller", internalType: "address", type: "address", indexed: true },
      { name: "requestIds", internalType: "uint256[]", type: "uint256[]", indexed: false },
    ],
    name: "TriggerBurnBatch",
  },
  {
    type: "event",
    anonymous: false,
    inputs: [
      { name: "caller", internalType: "address", type: "address", indexed: true },
      { name: "requestIds", internalType: "uint256[]", type: "uint256[]", indexed: false },
    ],
    name: "TriggerWithdrawal",
  },
  { type: "error", inputs: [], name: "InsufficientWithdrawal" },
  { type: "error", inputs: [], name: "InvalidRequestId" },
] as const
