<template>
  <div>
    <div class="sp-container">
      <sp-sign-in />
      <sp-bank-balances />
      <sp-token-send />
      <!-- this line is used by starport scaffolding # 4 -->
		<sp-type-form path="hirokryptor.xyzchain.xyzchain" type="pool" :fields="[ ['creator', 1, 'string'] , ['token1', 2, 'string'] , ['token2', 3, 'string'] , ['total1', 4, 'int32'] , ['total2', 5, 'int32'] ]" />
    </div>
  </div>
</template>

<script>
import * as sp from "@tendermint/vue";

export default {
  components: { ...sp },
};
</script>
