<script lang="ts">
import { Button, Dialog } from "$lib/elements";
</script>

<div class="min-h-screen bg-base-100 p-8">
  <div class="max-w-7xl mx-auto">
    <div class="space-y-16">
      <!-- Button Section -->
      <section>
        <div class="mb-8">
          <h2 class="text-2xl font-bold text-base-content mb-2">Button</h2>
          <p class="text-base-content/70">
            All button variants, styles, and sizes
          </p>
        </div>

        <div class="space-y-8">
          <!-- Variants -->
          <div class="card bg-base-200 p-6">
            <h3 class="text-lg font-semibold mb-4">Variants</h3>
            <div class="flex flex-wrap gap-4">
              <Button>Neutral</Button>
              <Button variant="primary">Primary</Button>
              <Button variant="secondary">Secondary</Button>
              <Button variant="accent">Accent</Button>
              <Button variant="info">Info</Button>
              <Button variant="success">Success</Button>
              <Button variant="warning">Warning</Button>
              <Button variant="error">Error</Button>
            </div>
          </div>

          <!-- Styles -->
          <div class="card bg-base-200 p-6">
            <h3 class="text-lg font-semibold mb-4">Styles</h3>
            <div class="flex flex-wrap gap-4">
              <Button variant="primary">Default</Button>
              <Button variant="primary" style="outline">Outline</Button>
              <Button variant="primary" style="soft">Soft</Button>
              <Button variant="primary" style="dash">Dash</Button>
              <Button variant="primary" style="ghost">Ghost</Button>
              <Button variant="primary" style="link">Link</Button>
            </div>
          </div>

          <!-- Sizes -->
          <div class="card bg-base-200 p-6">
            <h3 class="text-lg font-semibold mb-4">Sizes</h3>
            <div class="flex flex-wrap items-center gap-4">
              <Button variant="primary" size="xs">Extra Small</Button>
              <Button variant="primary" size="sm">Small</Button>
              <Button variant="primary" size="md">Medium</Button>
              <Button variant="primary" size="lg">Large</Button>
              <Button variant="primary" size="xl">Extra Large</Button>
            </div>
          </div>

          <!-- Modifiers -->
          <div class="card bg-base-200 p-6">
            <h3 class="text-lg font-semibold mb-4">Modifiers</h3>
            <div class="space-y-4">
              <div class="flex flex-wrap gap-4">
                <Button variant="primary">Normal</Button>
                <Button variant="primary" wide>Wide</Button>
                <Button variant="primary" square>SQ</Button>
                <Button variant="primary" circle>C</Button>
              </div>
              <div>
                <Button variant="primary" block>Block (Full Width)</Button>
              </div>
            </div>
          </div>

          <!-- States -->
          <div class="card bg-base-200 p-6">
            <h3 class="text-lg font-semibold mb-4">States</h3>
            <div class="flex flex-wrap gap-4">
              <Button variant="primary">Normal</Button>
              <Button variant="primary" active>Active</Button>
              <Button variant="primary" disabled>Disabled</Button>
            </div>
          </div>
        </div>
      </section>

      <!-- Dialog Section -->
      <section>
        <div class="mb-8">
          <h2 class="text-2xl font-bold text-base-content mb-2">Dialog</h2>
          <p class="text-base-content/70">
            Dialog modal components with different placements
          </p>
        </div>

        <div class="space-y-8">
          <!-- Basic Dialog -->
          <div class="card bg-base-200 p-6">
            <h3 class="text-lg font-semibold mb-4">Basic Dialog</h3>
            <div class="flex flex-wrap gap-4">
              <Dialog>
                {#snippet trigger(onclick)}
                  <Button variant="primary" {onclick}>Open Dialog</Button>
                {/snippet}
                <div class="modal-box">
                  <h3 class="text-lg font-bold">Hello!</h3>
                  <p class="py-4">Press ESC key or click outside to close</p>
                </div>
              </Dialog>
            </div>
          </div>

          <!-- Placement Variants -->
          <div class="card bg-base-200 p-6">
            <h3 class="text-lg font-semibold mb-4">Placement</h3>
            <div class="flex flex-wrap gap-4">
              <Dialog placement="top">
                {#snippet trigger(onclick)}
                  <Button variant="secondary" {onclick}>Top</Button>
                {/snippet}
                <div class="modal-box">
                  <h3 class="text-lg font-bold">Top Placement</h3>
                  <p class="py-4">This dialog is at the top</p>
                </div>
              </Dialog>

              <Dialog placement="middle">
                {#snippet trigger(onclick)}
                  <Button variant="accent" {onclick}>Middle</Button>
                {/snippet}
                <div class="modal-box">
                  <h3 class="text-lg font-bold">Middle Placement</h3>
                  <p class="py-4">This dialog is in the middle</p>
                </div>
              </Dialog>

              <Dialog placement="bottom">
                {#snippet trigger(onclick)}
                  <Button variant="info" {onclick}>Bottom</Button>
                {/snippet}
                <div class="modal-box">
                  <h3 class="text-lg font-bold">Bottom Placement</h3>
                  <p class="py-4">This dialog is at the bottom</p>
                </div>
              </Dialog>
            </div>
          </div>

          <!-- Dialog with Actions -->
          <div class="card bg-base-200 p-6">
            <h3 class="text-lg font-semibold mb-4">With Actions</h3>
            <div class="flex flex-wrap gap-4">
              <Dialog>
                {#snippet trigger(onclick)}
                  <Button variant="success" {onclick}>Confirm Action</Button>
                {/snippet}
                <div class="modal-box">
                  <h3 class="text-lg font-bold">Confirm Action</h3>
                  <p class="py-4">Are you sure you want to proceed?</p>
                  <div class="modal-action">
                    <Button variant="error" size="sm">Cancel</Button>
                    <Button variant="success" size="sm">Confirm</Button>
                  </div>
                </div>
              </Dialog>
            </div>
          </div>
        </div>
      </section>
    </div>
  </div>
</div>
