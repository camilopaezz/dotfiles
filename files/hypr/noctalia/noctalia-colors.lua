local primary   = "rgb(ebbf87)"
local surface   = "rgb(16130f)"
local secondary = "rgb(dac3a8)"
local error     = "rgb(ffb4ab)"

hl.config({
    general = {
        col = {
            active_border   = primary,
            inactive_border = surface,
        },
    },

    group = {
        col = {
            border_active        = secondary,
            border_inactive      = surface,
            border_locked_active = error,
            border_locked_inactive = surface,
        },

        groupbar = {
            col = {
                active              = secondary,
                inactive            = surface,
                locked_active       = error,
                locked_inactive     = surface,
            },
        },
    },
})
