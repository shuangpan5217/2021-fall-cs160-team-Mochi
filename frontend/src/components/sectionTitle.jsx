function SectionTitle({ title, subtitle }) {
    return (
        <>
            <p className="agenda section-header">
                {title} <span className="section-subheader">{subtitle}</span>
            </p>
        </>
    );
}

export default SectionTitle;
