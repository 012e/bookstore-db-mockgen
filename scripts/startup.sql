WITH invoice_totals AS (
    SELECT
        ii.invoice_id,
        SUM(
                ii.quantity * i.base_price *
                COALESCE(
                        (
                            SELECT
                                EXP(SUM(LN(1 + ip.percentage)))
                            FROM
                                item_prices ip
                            WHERE
                                ip.item_id = ii.item_id
                              AND ip.begin_date <= inv.created_at
                              AND ip.end_date >= inv.created_at
                              AND (1 + ip.percentage) > 0  -- Ensure positive values for LN()
                        ),
                        1  -- Default factor if no applicable percentages
                )
        )::numeric(12,2) AS total_price
    FROM
        invoices_items ii
            JOIN items i ON i.id = ii.item_id
            JOIN invoices inv ON inv.id = ii.invoice_id
    GROUP BY
        ii.invoice_id
)
UPDATE invoices inv
SET total = invoice_totals.total_price::money
FROM invoice_totals
WHERE inv.id = invoice_totals.invoice_id;